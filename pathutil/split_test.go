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

	"reflect"
)

func TestSplit(t *testing.T) {
	for _, test := range []struct {
		p     string
		parts []string
	}{
		{"a", []string{"a"}},
		{"a/b", []string{"a", "b"}},
		{"shanhu.io/smlvm", []string{"shanhu.io", "smlvm"}},
	} {
		parts, err := Split(test.p)
		if err != nil {
			t.Errorf("Split(%q) got error: %s", test.p, err)
			continue
		}

		if !reflect.DeepEqual(parts, test.parts) {
			t.Errorf(
				"Split(%q) want %v got %v",
				test.p, test.parts, parts,
			)
		}
	}
}

func TestSplitInvalidPath(t *testing.T) {
	for _, p := range []string{
		"",
		"/x",
		"x//y",
		"/",
		"a/b/c/",
	} {
		parts, err := Split(p)
		if err == nil {
			t.Errorf(
				"split path %q got parts %v, want error",
				p, parts,
			)
		}
	}
}
