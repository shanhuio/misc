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
	"encoding/json"

	"shanhu.io/misc/errcode"
)

// Header is the JWT header.
type Header struct {
	Alg   string `json:"alg"`
	Typ   string `json:"typ"`
	KeyID string `json:"kid,omitempty"` // Key ID.
}

func (h *Header) encode() (string, error) {
	return encodeSegment(h)
}

func decodeHeader(s string) (*Header, error) {
	bs, err := decodeSegmentBytes(s)
	if err != nil {
		return nil, err
	}
	h := new(Header)
	if err := json.Unmarshal(bs, h); err != nil {
		return nil, err
	}
	return h, nil
}

func checkHeader(got, want *Header) error {
	if got.KeyID != want.KeyID {
		return errcode.InvalidArgf("kid=%q, want %q", got.KeyID, want.KeyID)
	}
	if got.Alg != want.Alg {
		return errcode.InvalidArgf("alg=%q, want %q", got.Alg, want.Alg)
	}
	if got.Typ != want.Typ {
		return errcode.InvalidArgf("typ=%q, want %q", got.Typ, want.Typ)
	}
	return nil
}
