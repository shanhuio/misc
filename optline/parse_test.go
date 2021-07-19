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

package optline

import (
	"testing"
)

func TestParse(t *testing.T) {
	for _, test := range []struct {
		line string
		k, v string
	}{
		{`key: "value"`, "key", "value"},
		{"key: `value`\n\n", "key", "value"},
		{`key: "value 0"`, "key", "value 0"},
		{`k1: "value"`, "k1", "value"},
		{`_k: "value"`, "_k", "value"},
		{"k: `value`", "k", "value"},
	} {
		opt, err := Parse(test.line)
		if err != nil {
			t.Errorf("%q: unexpected error: %s", test.line, err)
			continue
		}
		if opt.Key != test.k {
			t.Errorf("%q: want key %q, got %q", test.line, test.k, opt.Key)
		}
		if opt.Value != test.v {
			t.Errorf("%q: want value %q, got %q", test.line, test.v, opt.Value)
		}
	}
}
