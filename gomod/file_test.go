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

package gomod

import (
	"testing"
)

func TestModulePath(t *testing.T) {
	for _, test := range []struct {
		content, mod string
	}{
		{`module shanhu.io/misc`, "shanhu.io/misc"},
		{"  module    shanhu.io/misc\t\t\t\n\nextra", "shanhu.io/misc"},
		{`module "shanhu.io/misc/v1"`, "shanhu.io/misc/v1"},
		{`module "shanhu.io/misc"`, "shanhu.io/misc"},
		{"// comment\nmodule x // tail\nnext line", "x"},
		{"module `x` // tail", "x"},
	} {
		got, err := modulePath([]byte(test.content))
		if err != nil {
			t.Errorf("modulePath(%q) got error: %s", test.content, err)
		} else if got != test.mod {
			t.Errorf(
				"modulePath(%q), want %q, got %q",
				test.content, test.mod, got,
			)
		}
	}
}
