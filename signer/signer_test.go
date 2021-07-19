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

package signer

import (
	"testing"

	"reflect"

	"shanhu.io/misc/rand"
)

func testSigner(t *testing.T, k []byte) {
	s := New(k)
	o := func(bs []byte) {
		signed := s.Sign(bs)
		ok, dat := s.Check(signed)
		if !ok {
			t.Error("check failed")
		} else if !reflect.DeepEqual(dat, bs) {
			t.Errorf("got %v, want %v", dat, bs)
		}

		h := s.SignHex(bs)
		ok, dat = s.CheckHex(h)
		if !ok {
			t.Error("check failed")
		} else if !reflect.DeepEqual(dat, bs) {
			t.Errorf("got %v, want %v", dat, bs)
		}
	}

	os := func(s string) { o([]byte(s)) }
	os("")
	os("something")
	os("            ")

	for i := 0; i < 5; i++ {
		o(rand.Bytes(10))
	}
}

func TestSigner(t *testing.T) {
	testSigner(t, nil)
	testSigner(t, []byte{})
	for i := 0; i < 3; i++ {
		testSigner(t, rand.Bytes(8))
	}
}
