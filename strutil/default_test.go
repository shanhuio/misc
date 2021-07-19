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
	"testing"
)

func TestDefault(t *testing.T) {
	want := "some string"
	for _, test := range []struct {
		input, def, want string
	}{
		{"", want, want},
		{want, "default", want},
	} {
		got := Default(test.input, test.def)
		if got != want {
			t.Errorf(
				"Default(%q, %q), want %q, got %q",
				test.input, test.def, test.want, got,
			)
		}
	}
}
