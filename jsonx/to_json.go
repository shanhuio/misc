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

package jsonx

import (
	"bytes"

	"shanhu.io/text/lexing"
)

func marshalValue(v value) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := encodeValue(buf, v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func marshalValueLexing(v value) ([]byte, []*lexing.Error) {
	bs, err := marshalValue(v)
	if err != nil {
		return nil, lexing.SingleErr(err)
	}
	return bs, nil
}

// ToJSON converts a JSONX stream into a JSON stream.
func ToJSON(input []byte) ([]byte, []*lexing.Error) {
	r := bytes.NewReader(input)
	p, _ := newParser("", r)
	v := parseValue(p)
	if errs := p.Errs(); errs != nil {
		return nil, errs
	}
	return marshalValueLexing(v)
}
