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
	"testing"
)

func TestCheckLabel(t *testing.T) {
	for _, name := range []string{
		"normal-label",
		"normal-long-label",
		"s",   // single letter
		"888", // number only
	} {
		if err := CheckLabel(name); err != nil {
			t.Errorf("%q is a valid label, but got error %q", name, err)
		}
	}

	for _, name := range []string{
		"",
		"A-very-very-very-very-very-very-very-very–very-very-long",
		"-label",
		"label-",
		"label--label",
		"label&label",
	} {
		if CheckLabel(name) == nil {
			t.Errorf("%q is an invalid label, but got no checking error", name)
		}
	}
}
