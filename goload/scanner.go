// Copyright (C) 2021  Shanhu Tech Inc.
//
// This program is free software: you can redistribute it and/or modify it
// under the terms of the GNU Affero General Public License as published by the
// Free Software Foundation, either version 3 of the License, or (at your
// option) any later version.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.  See the GNU Affero General Public License
// for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package goload

import (
	"fmt"
	"go/build"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"shanhu.io/misc/gomod"
	"shanhu.io/misc/pathutil"
)

// ScanOptions provides the options for scanning a Go language repository.
type ScanOptions struct {
	Context *build.Context

	// TestdataWhiteList provides a whitelist of "testdata" packages that are
	// valid ones and being imported.
	TestdataWhiteList map[string]bool

	// PkgBlackList is a list of packages that will be skipped. It will also
	// skip its sub packages.
	PkgBlackList map[string]bool
}

type scanner struct {
	path string
	ctx  *build.Context
	opts *ScanOptions
	res  *ScanResult

	gomod bool

	modRoot    string // import path for the mod root in non-mod mode
	modVerRoot string // import path for the mod root in mod enabled mode

	vendorStack    *vendorStack
	vendorLayers   map[string]*vendorLayer
	vendorScanning bool
}

func newScanner(p string, opts *ScanOptions) *scanner {
	if opts == nil {
		opts = new(ScanOptions)
	}

	ret := &scanner{
		path:         p,
		opts:         opts,
		vendorStack:  new(vendorStack),
		vendorLayers: make(map[string]*vendorLayer),
	}

	if opts.Context != nil {
		ret.ctx = opts.Context
	} else {
		ret.ctx = &build.Default
	}

	return ret
}

func (s *scanner) srcRoot() string {
	return filepath.Join(s.ctx.GOPATH, "src")
}

func (s *scanner) warning(dir string, err error) {
	s.res.Warnings = append(s.res.Warnings, &scanError{
		dir: dir,
		err: err,
	})
}

func (s *scanner) skipDir(dir *scanDir) bool {
	if inSet(s.opts.PkgBlackList, dir.path) {
		return true
	}
	base := dir.base
	if strings.HasPrefix(base, "_") || strings.HasPrefix(base, ".") {
		return true
	}
	if base == "testdata" && !inSet(s.opts.TestdataWhiteList, dir.path) {
		return true
	}
	return false
}

func (s *scanner) enterMod(p, vp string) {
	s.modRoot, s.modVerRoot = p, vp
}

func (s *scanner) exitMod() { s.modRoot, s.modVerRoot = "", "" }

func (s *scanner) handleDir(dir *scanDir) error {
	switch dir.base {
	case "vendor":
		s.res.HasVendor = true
	case "internal":
		s.res.HasInternal = true
	}

	mode := build.ImportComment

	if s.gomod {
		pkg, err := s.ctx.ImportDir(dir.dir, mode)
		if err != nil {
			if isNoGoError(err) {
				return nil
			}
			s.warning(dir.path, err)
			return nil
		}

		if len(pkg.GoFiles) == 0 && len(pkg.CgoFiles) == 0 {
			return nil
		}

		s.res.Pkgs[dir.path] = &Package{Build: pkg}
	} else if s.vendorScanning {
		if inSet(s.opts.PkgBlackList, "!"+dir.path) {
			return nil
		}

		// check if it is a package
		pkg, err := s.ctx.Import(dir.path, "", mode)
		if err != nil {
			if isNoGoError(err) {
				return nil
			}
			s.warning(dir.path, err)
			return nil
		}

		if len(pkg.GoFiles) == 0 && len(pkg.CgoFiles) == 0 {
			return nil
		}

		if dir.vendor != nil {
			dir.vendor.addPkg(dir.path)
		}

		s.res.Pkgs[dir.path] = &Package{Build: pkg}
	} else {
		pkg, found := s.res.Pkgs[dir.path]
		if !found {
			return nil
		}

		if s.modRoot != "" {
			pkg.ModRoot = s.modRoot
			pkg.ModVerRoot = s.modVerRoot
			modRel := pathutil.Relative(s.modRoot, dir.path)
			pkg.ModVerPath = path.Join(s.modVerRoot, modRel)
		}

		importMap := make(map[string]string)
		for _, imp := range pkg.Build.Imports {
			mapped, hit := s.vendorStack.mapImport(imp)
			if !hit {
				continue
			}
			importMap[imp] = mapped
		}

		if len(importMap) > 0 {
			pkg.ImportMap = importMap
		}
	}
	return nil
}

func (s *scanner) walk(dir *scanDir) error {
	info, err := os.Lstat(dir.dir)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return nil
	}
	if s.skipDir(dir) {
		return nil
	}

	f, err := os.Open(dir.dir)
	if err != nil {
		return err
	}

	defer f.Close()
	names, err := f.Readdirnames(-1)
	if err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}

	sort.Strings(names)

	if !s.vendorScanning && findInSorted(names, "vendor") {
		ly := s.vendorLayers[dir.path]
		if ly == nil {
			panic(fmt.Sprintf("vendor layer missing: %s", dir.path))
		}
		if len(ly.pkgs) > 0 {
			s.vendorStack.push(ly)
			defer s.vendorStack.pop()
		}
	}

	if !s.vendorScanning {
		if s.modRoot == "" && findInSorted(names, "go.mod") {
			p := filepath.Join(dir.dir, "go.mod")
			modFile, err := gomod.Parse(p)
			if err != nil {
				s.warning(dir.path, fmt.Errorf("parse go.mod: %s", err))
			} else if isValidModPath(dir.path, modFile.Name) {
				s.enterMod(dir.path, modFile.Name)
				defer s.exitMod()
			}
		}
	}

	if err := s.handleDir(dir); err != nil {
		return err
	}

	for _, name := range names {
		sub := dir.sub(name)

		if s.vendorScanning && name == "vendor" {
			ly := newVendorLayer(dir.path)
			s.vendorLayers[dir.path] = ly
			sub.vendor = ly
		}

		if err := s.walk(sub); err != nil {
			return err
		}
	}

	return nil
}
