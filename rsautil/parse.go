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
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"

	"golang.org/x/crypto/ssh"
)

var (
	errNotRSA = errors.New("public key is not an RSA key")
	errNoKey  = errors.New("no key")
)

// ParsePrivateKey parses the PEM encoded RSA private key.
func ParsePrivateKey(bs []byte) (*rsa.PrivateKey, error) {
	if len(bs) == 0 {
		return nil, errNoKey
	}

	b, _ := pem.Decode(bs)
	if b == nil {
		return nil, errors.New("pem decode failed")
	}

	if x509.IsEncryptedPEMBlock(b) {
		return nil, errors.New("key is encrypted")
	}

	return x509.ParsePKCS1PrivateKey(b.Bytes)
}

// ParsePrivateKeyFile parses the PEM encded RSA private key file.
func ParsePrivateKeyFile(f string) (*rsa.PrivateKey, error) {
	bs, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	return ParsePrivateKey(bs)
}

// ParsePublicKey parses a marshalled public key in SSH authorized key format.
func ParsePublicKey(bs []byte) (*rsa.PublicKey, error) {
	if len(bs) == 0 {
		return nil, errNoKey
	}

	k, _, _, _, err := ssh.ParseAuthorizedKey(bs)
	if err != nil {
		return nil, err
	}

	if k.Type() != "ssh-rsa" {
		return nil, errNotRSA
	}
	ck, ok := k.(ssh.CryptoPublicKey)
	if !ok {
		return nil, errNotRSA
	}

	ret, ok := ck.CryptoPublicKey().(*rsa.PublicKey)
	if !ok {
		return nil, errNotRSA
	}
	return ret, nil
}

// ParsePublicKeyFile parses a marshalled public key file in SSH authorized key
// file format.
func ParsePublicKeyFile(f string) (*rsa.PublicKey, error) {
	bs, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	return ParsePublicKey(bs)
}
