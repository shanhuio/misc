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
	"strconv"
	"strings"
)

func isValidModPath(p, modPath string) bool {
	if modPath == p {
		return false
	}

	prefix := p + "/v"
	if !strings.HasPrefix(modPath, prefix) {
		return false
	}

	ver := strings.TrimPrefix(modPath, prefix)
	if _, err := strconv.Atoi(ver); err != nil {
		return false
	}

	return true
}
