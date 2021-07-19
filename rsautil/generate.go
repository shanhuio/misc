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
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"golang.org/x/crypto/ssh"
)

func pemBlock(k *rsa.PrivateKey, pwd []byte) (*pem.Block, error) {
	const pemType = "RSA PRIVATE KEY"

	if pwd == nil {
		return &pem.Block{
			Type:  pemType,
			Bytes: x509.MarshalPKCS1PrivateKey(k),
		}, nil
	}

	return x509.EncryptPEMBlock(
		rand.Reader, pemType,
		x509.MarshalPKCS1PrivateKey(k),
		pwd, x509.PEMCipherDES,
	)
}

// GenerateKey generates a private/public key pair with the given passphrase.
// n is the bit size of the RSA key. When n is less than 0, 4096 is used.
func GenerateKey(passphrase []byte, n int) (pri, pub []byte, err error) {
	if n <= 0 {
		n = 4096
	}
	key, err := rsa.GenerateKey(rand.Reader, n)
	if err != nil {
		return nil, nil, err
	}

	b, err := pemBlock(key, passphrase)
	if err != nil {
		return nil, nil, err
	}

	priBuf := new(bytes.Buffer)
	if err := pem.Encode(priBuf, b); err != nil {
		return nil, nil, err
	}
	pubKey, err := ssh.NewPublicKey(&key.PublicKey)
	if err != nil {
		return nil, nil, err
	}
	pub = ssh.MarshalAuthorizedKey(pubKey)
	return priBuf.Bytes(), pub, nil
}
