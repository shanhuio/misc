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

package gosyntax

import (
	"go/scanner"
	"go/token"

	"shanhu.io/misc/syntax"
)

func tokType(t token.Token, lit string) string {
	switch t {
	case token.COMMENT:
		return "cm"
	case token.IDENT:
		if builtInFuncMap[lit] {
			return "bfunc"
		}
		if builtInTypeMap[lit] {
			return "btype"
		}
		return "ident"
	case token.INT, token.FLOAT, token.IMAG:
		return "num"
	case token.CHAR, token.STRING:
		return "str"
	}
	if t.IsOperator() {
		return "op"
	}
	if t.IsKeyword() {
		return "kw"
	}
	return "na"
}

// Tokens breaks a Go language program in a token stream.
func Tokens(bs []byte) []*syntax.Token {
	fset := token.NewFileSet()
	f := fset.AddFile("a.go", fset.Base(), len(bs))
	s := new(scanner.Scanner)
	var errs scanner.ErrorList
	s.Init(f, bs, errs.Add, scanner.ScanComments)

	var ret []*syntax.Token
	endPos := 0
	for {
		pos, t, lit := s.Scan()
		if t == token.SEMICOLON && lit != ";" {
			continue // this is actually white space
		}

		offset := f.Offset(pos)
		if offset > endPos {
			ret = append(ret, &syntax.Token{
				Type: "ws",
				Lit:  string(bs[endPos:offset]),
			})
		}

		ret = append(ret, &syntax.Token{
			Type: tokType(t, lit),
			Lit:  lit,
		})

		if t == token.EOF {
			break
		}
		endPos = offset + len(lit)
	}

	return ret
}

// HTML renders a G language file into HTML file.
func HTML(bs []byte) string {
	toks := Tokens(bs)
	return syntax.RenderHTML(toks)
}
