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
	"go/build"

	"golang.org/x/tools/go/loader"
)

// Program is a loaded program from a set of packages.
type Program struct {
	*loader.Program

	Pkgs []string
}

// LoadPkgs loads a list of package from source files into a loader program.
func LoadPkgs(pkgs []string) (*Program, error) {
	b := build.Default
	b.CgoEnabled = false

	conf := loader.Config{Build: &b}
	for _, p := range pkgs {
		conf.Import(p)
	}

	ret, e := conf.Load()
	if e != nil {
		return nil, e
	}

	pkgsCopied := make([]string, len(pkgs))
	copy(pkgsCopied, pkgs)

	return &Program{Program: ret, Pkgs: pkgsCopied}, e
}

// LoadPkg loads a set of pacakges from source files into a loader program.
// It just calls LoadPkgs with the parameter list
func LoadPkg(pkgs ...string) (*Program, error) {
	return LoadPkgs(pkgs)
}
