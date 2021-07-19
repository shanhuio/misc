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

package rsautil

import (
	"bytes"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"

	"golang.org/x/crypto/ssh"
)

func keyHashStr(h []byte) string {
	return base64.RawURLEncoding.EncodeToString(h)
}

// PublicKeyHash returns the public key hash of a key.
func PublicKeyHash(k *rsa.PublicKey) ([]byte, error) {
	sshPub, err := ssh.NewPublicKey(k)
	if err != nil {
		return nil, err
	}

	wire := bytes.TrimSpace(ssh.MarshalAuthorizedKey(sshPub))
	h := sha256.Sum256(wire)
	return h[:], nil
}

// PublicKeyHashString returns the public key hash string of a key.
func PublicKeyHashString(k *rsa.PublicKey) (string, error) {
	h, err := PublicKeyHash(k)
	if err != nil {
		return "", err
	}
	return keyHashStr(h), nil
}
