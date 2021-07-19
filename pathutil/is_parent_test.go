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
	"testing"
)

func TestIsParent(t *testing.T) {
	for _, test := range []struct {
		short, long string
		want        bool
	}{
		{"a", "b", false},
		{"a", "a", true},
		{"a/b", "a", false},
		{"a", "a/b", true},
		{"a", "ab", false},
		{"a", "ab/c", false},
		{"a/b", "a/b/c", true},
	} {
		got := IsParent(test.short, test.long)
		if got != test.want {
			t.Errorf(
				"IsParent(%q, %q), want %v, got %v",
				test.short, test.long, test.want, got,
			)
		}
	}
}
