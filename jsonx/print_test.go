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

package jsonx

import (
	"testing"

	"bytes"
	"encoding/json"
)

func TestMarshal_loopback(t *testing.T) {
	for _, obj := range []interface{}{
		"something",
		1.234,
		1234,
		nil,
		struct{ A, B string }{A: "a", B: "b"},
		map[string]string{
			"a.com": "a:8888",
			"b.com": "b:7777",
		},
		[]int{1, 2, 3},
	} {
		want, err := json.Marshal(obj)
		if err != nil {
			t.Fatalf("marshal %v: %v", obj, err)
		}

		x, err := Marshal(obj)
		if err != nil {
			t.Errorf("format %v: %v", obj, err)
			continue
		}

		var box interface{}
		if err := Unmarshal(x, &box); err != nil {
			t.Errorf("unmarshal %q: %v", x, err)
			continue
		}

		got, err := json.Marshal(box)
		if err != nil {
			t.Fatalf("marshal jsonx-gen %v: %v", obj, err)
		}

		if !bytes.Equal(want, got) {
			t.Errorf(
				"format test failed %v: got %q, want %q",
				obj, got, want,
			)
		}
	}
}
