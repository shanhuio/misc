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

	"reflect"
)

func TestMakeSet(t *testing.T) {
	for _, test := range []struct {
		list []string
		want map[string]bool
	}{
		{nil, map[string]bool{}},
		{[]string{}, map[string]bool{}},
		{[]string{"a", "B"}, map[string]bool{"a": true, "B": true}},
		{[]string{"a", "a"}, map[string]bool{"a": true}},
	} {
		got := MakeSet(test.list)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf(
				"MakeSet(%v): got %v, want %v",
				test.list, got, test.want,
			)
		}
	}
}
