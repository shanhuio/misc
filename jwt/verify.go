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
	"strings"

	"shanhu.io/misc/errcode"
)

// Verifier verifies the token.
type Verifier interface {
	Verify(h *Header, data, sig []byte) error
}

// Token is a parsed JWT token.
type Token struct {
	Header    *Header
	ClaimSet  *ClaimSet
	Signature []byte
}

// DecodeAndVerify decodes and verifies a token.
func DecodeAndVerify(token string, v Verifier) (*Token, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, errcode.InvalidArgf(
			"invalid token: %d parts", len(parts),
		)
	}

	h, c, sig := parts[0], parts[1], parts[2]
	header := new(Header)
	if err := decodeSegment(h, header); err != nil {
		return nil, errcode.InvalidArgf("decode header: %s", err)
	}

	payload := []byte(token[:len(h)+1+len(c)])
	sigBytes, err := decodeSegmentBytes(sig)
	if err != nil {
		return nil, errcode.InvalidArgf("decode signature: %s", err)
	}

	if v != nil {
		if err := v.Verify(header, payload, sigBytes); err != nil {
			return nil, errcode.Annotate(err, "verify signature")
		}
	}

	claims, err := decodeClaimSet(c)
	if err != nil {
		return nil, errcode.InvalidArgf("decode claims: %s", err)
	}

	return &Token{
		Header:    header,
		ClaimSet:  claims,
		Signature: sigBytes,
	}, nil
}

// Decode decodes the token without verifying it.
func Decode(token string) (*Token, error) {
	return DecodeAndVerify(token, nil)
}
