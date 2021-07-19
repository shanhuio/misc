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

package strutil

import (
	"sort"
)

// MakeSet converts a list of strings to a set of strings.
func MakeSet(lst []string) map[string]bool {
	ret := make(map[string]bool)
	for _, s := range lst {
		ret[s] = true
	}
	return ret
}

// SortedList returns the sorted list of a set of strings.
func SortedList(set map[string]bool) []string {
	var ret []string
	for s := range set {
		ret = append(ret, s)
	}
	sort.Strings(ret)
	return ret
}
