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
	"fmt"
	"io"
	"log"
	"os"
)

const formatIndent = "  "

// Fprint pretty prints a JSON data blob into a writer.
func Fprint(w io.Writer, v interface{}) error {
	bs, err := json.MarshalIndent(v, "", formatIndent)
	if err != nil {
		return err
	}
	if _, err := w.Write(bs); err != nil {
		return err
	}
	_, err = fmt.Fprintln(w)
	return err
}

// Print pretty prints a JSON data blob into stdout.
func Print(v interface{}) {
	if err := Fprint(os.Stdout, v); err != nil {
		log.Println(err)
	}
}

// Format pretty formats JSON data bytes.
func Format(bs []byte) ([]byte, error) {
	out := new(bytes.Buffer)
	if err := json.Indent(out, bs, "", formatIndent); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}
