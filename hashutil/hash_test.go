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

package hashutil

import (
	"testing"

	"strings"
)

func TestHash(t *testing.T) {
	m := make(map[string]bool)
	addHash := func(h string) {
		if m[h] {
			t.Fatalf("hash conflict: %s", h)
		}
		m[h] = true
	}

	addHash(Hash(nil))
	addHash(HashStr("a"))
	addHash(HashStr("A"))
	addHash(HashStr("A "))
	addHash(HashStr("Hello"))

	const s = "something"
	h1 := HashStr(s)
	h2, err := HashReader(strings.NewReader(s))
	if err != nil {
		t.Fatal(err)
	}
	if h1 != h2 {
		t.Errorf("HashStr(%q) != HashReader(%q)", s, s)
	}
}

func TestHashFile(t *testing.T) {
	got, err := HashFile("testdata/testfile")
	if err != nil {
		t.Fatal(err)
	}

	want := HashStr("something\n")
	if want != got {
		t.Errorf("HashFile want %q, got %q", want, got)
	}
}
