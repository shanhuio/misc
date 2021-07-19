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

func TestValidPath(t *testing.T) {
	o := func(p string) {
		if !ValidPath(p) {
			t.Errorf("%q shoud be valid", p)
		}
	}

	e := func(p string) {
		if ValidPath(p) {
			t.Errorf("%q should be invalid", p)
		}
	}

	o("/")
	o("/asdf")
	o("/valentines_day")
	o("/thank/you")
	o("/3307")
	o("/c323/b75/53_df_")

	e("")
	e("/Hello")
	e("//")
	e("/as/")
	e("/as//of")
	e("/asdf-er")
	e("/  ")
	e("asdf")
	e("/2014-01-18")
}
