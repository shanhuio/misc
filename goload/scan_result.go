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
)

// Package is a package in the scan result.
type Package struct {
	Build     *build.Package
	ImportMap map[string]string // import remapping due to vendoring

	ModRoot    string // module root import path, not including version
	ModVerRoot string // module root import path, including version
	ModVerPath string // alias import path when module is enabled
}

// ScanResult has the scanning result
type ScanResult struct {
	Repo        string
	Pkgs        map[string]*Package
	HasVendor   bool
	HasInternal bool
	Warnings    []error
}

func newScanResult(repo string) *ScanResult {
	return &ScanResult{
		Repo: repo,
		Pkgs: make(map[string]*Package),
	}
}
