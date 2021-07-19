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
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strconv"

	"shanhu.io/text/fmtutil"
)

type printer struct {
	p *fmtutil.Printer
}

func newPrinter(w io.Writer) *printer {
	p := fmtutil.NewPrinter(w)
	return &printer{p: p}
}

func (p *printer) writeString(s string) {
	io.WriteString(p.p, strconv.Quote(s))
}

func (p *printer) write(v interface{}) {
	switch v := v.(type) {
	case bool:
		io.WriteString(p.p, strconv.FormatBool(v))
	case float64:
		s := strconv.FormatFloat(v, 'g', -1, 64)
		io.WriteString(p.p, s)
	case string:
		p.writeString(v)
	case []interface{}:
		p.writeArray(v)
	case map[string]interface{}:
		p.writeObject(v)
	case nil:
		io.WriteString(p.p, "null")
	}
}

func (p *printer) writeArrayItems(array []interface{}) {
	p.p.Tab()
	defer p.p.ShiftTab()

	for _, item := range array {
		p.write(item)
		io.WriteString(p.p, ",\n")
	}
}

func (p *printer) writeArray(array []interface{}) {
	if len(array) == 0 {
		io.WriteString(p.p, "[]")
		return
	}
	io.WriteString(p.p, "[\n")
	p.writeArrayItems(array)
	io.WriteString(p.p, "]")
}

func isIdent(s string) bool {
	if s == "" {
		return false
	}
	if _, is := keywords[s]; is {
		return false
	}

	for i, r := range s {
		if r == '_' {
			continue
		}
		if r >= 'a' && r <= 'z' {
			continue
		}
		if r >= 'A' && r <= 'Z' {
			continue
		}
		if r >= '0' && r <= '9' && i > 0 {
			continue
		}
		return false
	}
	return true
}

func (p *printer) writeObjectItems(obj map[string]interface{}) {
	var keys []string
	identKeys := true
	for k := range obj {
		keys = append(keys, k)
		if !isIdent(k) {
			identKeys = false
		}
	}
	sort.Strings(keys)

	p.p.Tab()
	defer p.p.ShiftTab()

	for _, k := range keys {
		if identKeys {
			io.WriteString(p.p, k)
		} else {
			p.writeString(k)
		}
		io.WriteString(p.p, ": ")
		p.write(obj[k])
		io.WriteString(p.p, ",\n")
	}
}

func (p *printer) writeObject(obj map[string]interface{}) {
	if len(obj) == 0 {
		io.WriteString(p.p, "{}")
		return
	}

	io.WriteString(p.p, "{\n")
	p.writeObjectItems(obj)
	io.WriteString(p.p, "}")
}

func (p *printer) err() error { return p.p.Err() }

// Fprint formats v in JSONX and prints it into w.
func Fprint(w io.Writer, v interface{}) error {
	bs, err := json.Marshal(v)
	if err != nil {
		return err
	}

	var g interface{}
	if err := json.Unmarshal(bs, &g); err != nil {
		return err
	}

	p := newPrinter(w)
	p.write(g)
	if err := p.err(); err != nil {
		return err
	}
	_, err = io.WriteString(w, "\n")
	return err
}

// Print formats v in JSONX and prints it into stdout.
func Print(v interface{}) error {
	return Fprint(os.Stdout, v)
}

// Sprint formats v in JSONX and returns the formatted string.
func Sprint(v interface{}) (string, error) {
	buf := new(bytes.Buffer)
	if err := Fprint(buf, v); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// Marshal formats v in JSONX.
func Marshal(v interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := Fprint(buf, v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// WriteFile writes a JSONX object into a file.
func WriteFile(p string, v interface{}) error {
	bs, err := Marshal(v)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(p, bs, 0644)
}
