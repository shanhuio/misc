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

package markdown

import (
	"bytes"
	"fmt"

	"golang.org/x/net/html"
)

// Compiler compiles a source string into a series of bytes.
type Compiler interface {
	Compile(src string) ([]byte, error)
}

// Compile goes through the given HTML and compiles the smallrepo code plugins
// using the given compiler.
func Compile(src []byte, c Compiler) ([]byte, error) {
	r := bytes.NewReader(src)
	doc, err := html.Parse(r)
	if err != nil {
		return nil, fmt.Errorf("html parse: %s", err)
	}

	w := new(bytes.Buffer)
	if err := html.Render(w, doc); err != nil {
		return nil, fmt.Errorf("html render: %s", err)
	}

	return w.Bytes(), nil
}
