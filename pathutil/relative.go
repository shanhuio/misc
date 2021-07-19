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

package pathutil

import (
	"strings"
)

// DotRelative returns the relative path of full to base.
// The return value starts with a dot if base is a parent of full.
// It returns full unchanged if base is not a parent of full.
func DotRelative(base, full string) string {
	if base == full {
		return "."
	}
	if strings.HasPrefix(full, base+"/") {
		return "./" + strings.TrimPrefix(full, base+"/")
	}
	return full
}

// Relative returns the relative path of full to base.
// It returns the relative path if base is a parent of full.
// It returns a single dot if base and full are the same.
// It retunrs an empty string if base is not a parent of full.
func Relative(base, full string) string {
	if base == full {
		return "."
	}
	if strings.HasPrefix(full, base+"/") {
		return strings.TrimPrefix(full, base+"/")
	}
	return ""
}
