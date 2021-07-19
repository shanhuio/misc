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
	"encoding/base64"
	"encoding/json"
)

func encodeSegmentBytes(bs []byte) string {
	return base64.RawURLEncoding.EncodeToString(bs)
}

func decodeSegmentBytes(s string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(s)
}

func encodeSegment(v interface{}) (string, error) {
	bs, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return encodeSegmentBytes(bs), nil
}

func decodeSegment(s string, v interface{}) error {
	bs, err := decodeSegmentBytes(s)
	if err != nil {
		return err
	}
	return json.Unmarshal(bs, v)
}
