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
	"strings"

	"golang.org/x/tools/go/buildutil"
)

func skipPkg(p string) bool {
	if p == "C" {
		return true
	}
	if p == "lib9" {
		return true
	}
	if strings.HasPrefix(p, "lib9/") {
		return true
	}
	if p == "cmd" {
		return true
	}
	if strings.HasPrefix(p, "cmd/") {
		return true
	}
	return false
}

// ListStdPkgs returns a list of all packages in
// the golang standard library with the specific GOPATH.
func ListStdPkgs() []string {
	c := build.Default
	c.CgoEnabled = false
	c.GOPATH = ""

	pkgs := buildutil.AllPackages(&c)

	var ret []string
	for _, p := range pkgs {
		if skipPkg(p) {
			continue
		}

		if _, err := build.Import(p, "", 0); err != nil {
			continue // ignore error packges
		}

		ret = append(ret, p)
	}
	return ret
}
