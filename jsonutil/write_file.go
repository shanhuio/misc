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

package jsonutil

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

// WriteFile marshals a JSON object and writes it into a file.
func WriteFile(file string, obj interface{}) error {
	bs, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(file, bs, 0644)
}

// WriteFileReadable marshals a JSON object with indents and writes it into a
// file.
func WriteFileReadable(f string, v interface{}) error {
	buf := new(bytes.Buffer)
	bs, err := json.MarshalIndent(v, "", formatIndent)
	if err != nil {
		return err
	}
	buf.Write(bs)
	buf.Write([]byte("\n"))

	return ioutil.WriteFile(f, buf.Bytes(), 0644)
}
