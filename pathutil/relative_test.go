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

func TestRelative(t *testing.T) {
	for _, test := range []struct {
		base, full string
		want       string
	}{
		{"a", "b", ""},
		{"a", "a/b", "b"},
		{"a", "a", "."},
		{"a", "a/b/c", "b/c"},
		{"a", "ab/c", ""},
	} {
		got := Relative(test.base, test.full)
		if got != test.want {
			t.Errorf(
				"Relative(%q, %q), want %q, got %q",
				test.base, test.full, test.want, got,
			)
		}
	}
}

func TestDotRelative(t *testing.T) {
	for _, test := range []struct {
		base, full string
		want       string
	}{
		{"a", "b", "b"},
		{"a", "a/b", "./b"},
		{"a", "a", "."},
		{"a", "a/b/c", "./b/c"},
		{"a", "ab/c", "ab/c"},
	} {
		got := DotRelative(test.base, test.full)
		if got != test.want {
			t.Errorf(
				"DotRelative(%q, %q), want %q, got %q",
				test.base, test.full, test.want, got,
			)
		}
	}
}
