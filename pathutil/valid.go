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

	"shanhu.io/text/lexing"
)

// ValidPathRune checks if r is a valid rune in the path.
// Valid runes contains a-z, A-Z, 0-9, '_' and '.'
func ValidPathRune(r rune) bool {
	if r == '_' || r == '.' {
		return true
	}
	if r >= 'a' && r <= 'z' {
		return true
	}
	return lexing.IsDigit(r)
}

// ValidPath checks if p is a valid absolute path
func ValidPath(p string) bool {
	if !strings.HasPrefix(p, "/") {
		return false
	}
	if p == "/" {
		return true
	}

	p = strings.TrimPrefix(p, "/")
	subs := strings.Split(p, "/")
	for _, s := range subs {
		if s == "" {
			return false
		}

		for _, r := range s {
			if !ValidPathRune(r) {
				return false
			}
		}
	}
	return true
}
