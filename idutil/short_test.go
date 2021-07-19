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

package idutil

import (
	"testing"
)

func TestShortId(t *testing.T) {
	for _, test := range []struct {
		id   string
		want string
	}{
		{"", ""},
		{"123", "123"},
		{"1234567", "1234567"},
		{"12345678", "1234567"},
		{"1234567890", "1234567"},
		{"\\\\", ""},
		{"汉字？？", ""},
	} {
		got := Short(test.id)
		if got != test.want {
			t.Errorf(
				"Short id string for %q: got %q, want %q",
				test.id, got, test.want,
			)
		}
	}
}
