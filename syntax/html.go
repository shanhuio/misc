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

package syntax

import (
	"bytes"
	"fmt"
	"html"
)

func runeHTML(r rune) string {
	switch r {
	case '\t':
		return "&nbsp;&nbsp;&nbsp;&nbsp;"
	case ' ':
		return "&nbsp;"
	case '\n':
		return "<br>\n"
	}
	return html.EscapeString(string(r))
}

func writeToken(buf *bytes.Buffer, tok *Token) {
	fmt.Fprintf(buf, `<span class="%s">`, tok.Type)
	for _, r := range tok.Lit {
		fmt.Fprint(buf, runeHTML(r))
	}
	fmt.Fprint(buf, "</span>")
}

// RenderHTML renders a token series into a HTML file.
func RenderHTML(toks []*Token) string {
	buf := new(bytes.Buffer)
	for _, t := range toks {
		writeToken(buf, t)
	}
	return buf.String()
}
