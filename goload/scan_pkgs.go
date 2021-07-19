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
	"path"
	"path/filepath"
	"sort"
)

// ScanModPkgs scans all packages in a module.
func ScanModPkgs(mod, dir string, opts *ScanOptions) (*ScanResult, error) {
	s := newScanner(mod, opts)
	s.gomod = true
	s.res = newScanResult(mod)
	d := &scanDir{dir: dir, path: mod}

	s.vendorScanning = false
	if err := s.walk(d); err != nil {
		return nil, err
	}
	return s.res, nil
}

// ScanPkgs scans all packages under a package path.
func ScanPkgs(p string, opts *ScanOptions) (*ScanResult, error) {
	s := newScanner(p, opts)

	// First check if the folder can be found.
	s.res = newScanResult(p)
	dir := &scanDir{
		dir:  filepath.Join(s.srcRoot(), filepath.ToSlash(p)),
		path: p,
		base: path.Base(p),
	}

	for _, scanning := range []bool{true, false} {
		s.vendorScanning = scanning
		if err := s.walk(dir); err != nil {
			return nil, err
		}
	}

	return s.res, nil
}

// ListPkgs list all packages under a package path.
func ListPkgs(p string) ([]string, error) {
	res, err := ScanPkgs(p, nil)
	if err != nil {
		return nil, err
	}

	var lst []string
	for pkg := range res.Pkgs {
		lst = append(lst, pkg)
	}
	sort.Strings(lst)
	return lst, nil
}
