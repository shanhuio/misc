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
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	"shanhu.io/misc/rand"
)

// Signer is a signer that contains a secrect key.
type Signer struct {
	key []byte
}

// New creates a signing pen.
func New(key []byte) *Signer {
	if key == nil {
		key = rand.Bytes(32)
	}
	return &Signer{key: key}
}

func (s *Signer) hash(dat []byte) []byte {
	m := hmac.New(sha256.New, s.key)
	m.Write(dat)
	return m.Sum(nil)
}

// Sign signs a blob and returns the combination of the data and the signature.
func (s *Signer) Sign(dat []byte) []byte {
	buf := new(bytes.Buffer)
	buf.Write(dat)

	h := s.hash(buf.Bytes())
	buf.Write(h)

	return buf.Bytes()
}

// SignJSON signs a JSON marshalable blob and returns the combination of the
// data and the signature.
func (s *Signer) SignJSON(dat interface{}) ([]byte, error) {
	bs, err := json.Marshal(dat)
	if err != nil {
		return nil, err
	}

	return s.Sign(bs), nil
}

// SignHex signs a blob and returns the data along with the signature in a hex
// string.
func (s *Signer) SignHex(dat []byte) string {
	return hex.EncodeToString(s.Sign(dat))
}

// SignHexJSON signs a JSON marshalable blob and returns the data along with
// the signature in a hex string.
func (s *Signer) SignHexJSON(dat interface{}) (string, error) {
	bs, err := s.SignJSON(dat)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bs), nil
}

// Check verifies if the signed blob is valid. If it is valid, it returns the
// original data that is protected by the signature.
func (s *Signer) Check(bs []byte) (bool, []byte) {
	n := len(bs)
	if n < sha256.Size {
		return false, nil
	}

	dat := bs[:n-sha256.Size]
	hashGot := bs[n-sha256.Size:]
	hashWant := s.hash(dat)
	if !hmac.Equal(hashGot, hashWant) {
		return false, nil
	}
	return true, dat
}

// CheckJSON verifies if the signed blob is valid, and if it is, unmarshals
// the original data into dat.
func (s *Signer) CheckJSON(bs []byte, dat interface{}) (bool, error) {
	ok, checked := s.Check(bs)
	if !ok {
		return false, nil
	}
	return true, json.Unmarshal(checked, dat)
}

// CheckHexJSON verifies if the signed blob is valid, and if it is, unmarshals
// the original data into dat.
func (s *Signer) CheckHexJSON(str string, dat interface{}) (bool, error) {
	ok, bs := s.CheckHex(str)
	if !ok {
		return false, nil
	}
	return true, json.Unmarshal(bs, dat)
}

// CheckHex verifies if the signed blob is valid, and if it is, returns the
// original data that is protected by the signature.
func (s *Signer) CheckHex(str string) (bool, []byte) {
	bs, err := hex.DecodeString(str)
	if err != nil {
		return false, nil
	}
	return s.Check(bs)
}
