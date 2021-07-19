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

package nameutil

import (
	"shanhu.io/misc/errcode"
)

// CheckLabel checks whether a string can be safely used as a
// sub-domain name.
func CheckLabel(s string) error {
	if len(s) == 0 {
		return errcode.InvalidArgf("empty name")
	}
	if len(s) > 50 {
		return errcode.InvalidArgf("name too long: %q", s)
	}

	if s[0] == '-' {
		return errcode.InvalidArgf("%q starts with hypen", s)
	}
	if s[len(s)-1] == '-' {
		return errcode.InvalidArgf("%q ends with hypen", s)
	}
	lastHyphen := false
	for _, r := range s {
		if r == '-' {
			if lastHyphen {
				return errcode.InvalidArgf("%q has continous hyphen", s)
			}
			lastHyphen = true
			continue
		} else {
			lastHyphen = false
		}
		if r >= '0' && r <= '9' {
			continue
		}
		if r >= 'a' && r <= 'z' {
			continue
		}
		return errcode.InvalidArgf("%q contain invalid char: %q", s, r)
	}
	return nil
}
