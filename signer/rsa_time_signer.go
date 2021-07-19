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
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"time"

	"shanhu.io/misc/rsautil"
)

// SignedRSABlock is a signed RSA block.
type SignedRSABlock struct {
	Data []byte
	Hash []byte
	Sig  []byte

	KeyID string `json:",omitempty"`
}

// RSATimeSigner signes the current time, or checks if a signed time
// is within a time window of the current time reading.
type RSATimeSigner struct {
	k      *rsa.PublicKey
	window time.Duration

	TimeFunc func() time.Time
}

// NewRSATimeSigner creates a new time signer that uses an RSA key.
func NewRSATimeSigner(k *rsa.PublicKey, w time.Duration) *RSATimeSigner {
	if w < 0 {
		w = -w
	}

	return &RSATimeSigner{
		k:      k,
		window: w,
	}
}

func rsaSignTime(k *rsa.PrivateKey, t time.Time) (*SignedRSABlock, error) {
	buf := make([]byte, timestampLen)
	binary.LittleEndian.PutUint64(buf, uint64(t.UnixNano()))
	hash := sha256.Sum256(buf)
	sig, err := rsa.SignPKCS1v15(rand.Reader, k, crypto.SHA256, hash[:])
	if err != nil {
		return nil, fmt.Errorf("sign blob: %s", err)
	}
	keyHash, err := rsautil.PublicKeyHashString(&k.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("make key hash: %s", err)
	}

	return &SignedRSABlock{
		Data:  buf,
		Hash:  hash[:],
		Sig:   sig,
		KeyID: keyHash,
	}, nil
}

// RSASignTime signes the current time with the given RSA key.
func RSASignTime(k *rsa.PrivateKey) (*SignedRSABlock, error) {
	return rsaSignTime(k, time.Now())
}

// Check checks if the timestamp is with in the time window.
func (s *RSATimeSigner) Check(b *SignedRSABlock) error {
	if len(b.Data) < 8 {
		return fmt.Errorf("data too short to have a timestamp")
	}
	t := time.Unix(0, int64(binary.LittleEndian.Uint64(b.Data)))
	timeNow := now(s.TimeFunc)
	if !inWindow(t, timeNow, s.window) {
		return fmt.Errorf("time out of window")
	}

	hash := sha256.Sum256(b.Data)
	if !bytes.Equal(hash[:], b.Hash) {
		return fmt.Errorf("hash incorrect")
	}
	return rsa.VerifyPKCS1v15(s.k, crypto.SHA256, b.Hash, b.Sig)
}

// CheckRSATimeSignature checks if the signed RSA block is signed with the
// given key, and with in the time window.
func CheckRSATimeSignature(
	b *SignedRSABlock, k *rsa.PublicKey, w time.Duration,
) error {
	s := NewRSATimeSigner(k, w)
	return s.Check(b)
}
