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

package jwt

import (
	"testing"

	"time"

	"shanhu.io/misc/rand"
)

func TestHS256(t *testing.T) {
	key := rand.Bytes(32) // 256 bits
	h := NewHS256(key, "")
	now := time.Now()
	c := &ClaimSet{
		Iss: "shanhu.io",
		Aud: "nextcloud",
		Iat: now.Unix(),
		Exp: now.Add(time.Hour).Unix(),
		Sub: "h8liu",
	}

	tokStr, err := EncodeAndSign(c, h)
	if err != nil {
		t.Fatal("encode: ", err)
	}
	t.Log(tokStr)

	tok, err := DecodeAndVerify(tokStr, h)
	if err != nil {
		t.Fatal("decode: ", err)
	}

	if got, want := tok.ClaimSet.Iss, c.Iss; got != want {
		t.Errorf("got issuer %q, want %q", got, want)
	}

	if _, err := CheckTime(tok.ClaimSet, now); err != nil {
		t.Errorf("check time: %s", err)
	}
}
