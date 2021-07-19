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

// Package gomod provides simple go.mod file parsing.
//
// TODO(h8liu): simply call golang.org/x/mod/modfile
package gomod

import (
	"bufio"
	"bytes"
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

// File is a parsed go.mod file.
type File struct {
	Name string
}

// Parse parses a go.mod file.
func Parse(f string) (*File, error) {
	bs, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	name, err := modulePath(bs)
	if err != nil {
		return nil, err
	}
	return &File{Name: name}, nil
}

var errInvalidModFile = errors.New("invalid go.mod file")

// modulePath returns the module path from the gomod file text.
// If it cannot find a module path, it returns an empty string.
// It is tolerant of unrelated problems in the go.mod file.
func modulePath(bs []byte) (string, error) {
	s := bufio.NewScanner(bytes.NewReader(bs))

	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}
		if !strings.HasPrefix(line, "module") {
			continue
		}

		line = strings.TrimSpace(strings.TrimPrefix(line, "module"))

		// TODO: this is incorrect for quoted module path
		if pos := strings.Index(line, "//"); pos >= 0 {
			line = strings.TrimSpace(line[:pos])
			if line == "" {
				return "", errInvalidModFile
			}
		}

		if line == "" {
			return "", errInvalidModFile
		}
		if line[0] == '"' || line[0] == '`' {
			p, err := strconv.Unquote(line)
			if err != nil || p == "" {
				return "", errInvalidModFile
			}
			return p, nil
		}

		return line, nil
	}

	return "", errInvalidModFile
}
