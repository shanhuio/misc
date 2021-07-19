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

// Package strtoken parses a line into string tokens.  Tokens are separated by
// white spaces, and can be double quoted for tokens that contains white
// spaces.
package strtoken

import (
	"io"
	"strconv"
	"strings"

	"shanhu.io/text/lexing"
)

func lexShell(x *lexing.Lexer) *lexing.Token {
	r := x.Rune()
	if x.IsWhite(r) {
		panic("incorrect token start")
	}

	switch r {
	case '"':
		return lexing.LexString(x, str, '"')
	}

	if isBareRune(r) {
		return lexBare(x)
	}

	x.Errorf("illegal char %q", r)
	x.Next()
	return x.MakeToken(lexing.Illegal)
}

func newLexer(file string, r io.Reader) *lexing.Lexer {
	return lexing.MakeLexer(file, r, lexShell)
}

// Parse parses a line into a sequence of tokens.
// A token
func Parse(line string) ([]string, []*lexing.Error) {
	r := strings.NewReader(line)
	x := newLexer("", r)
	toks := lexing.TokenAll(x)
	if errs := x.Errs(); errs != nil {
		return nil, errs
	}

	var ret []string
	var errs []*lexing.Error
	for _, t := range toks {
		switch t.Type {
		case bare:
			ret = append(ret, t.Lit)
		case str:
			v, err := strconv.Unquote(t.Lit)
			if err != nil {
				errs = append(errs, &lexing.Error{
					Pos:  t.Pos,
					Err:  err,
					Code: "shellarg.invalidStr",
				})
			} else {
				ret = append(ret, v)
			}
		}
	}

	if errs != nil {
		return nil, errs
	}
	return ret, nil
}
