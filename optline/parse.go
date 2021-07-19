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

package optline

import (
	"fmt"
	"shanhu.io/text/lexing"
	"strconv"
	"strings"
)

// Option is a single key value option.
type Option struct {
	Key   string
	Value string
}

const (
	tokIdent = iota
	tokCollon
	tokString
)

func tokStr(tok int) string {
	switch tok {
	case tokIdent:
		return "ident"
	case tokCollon:
		return "`:`"
	case tokString:
		return "string"
	default:
		return fmt.Sprintf("?%d", tok)
	}
}

func lexLine(x *lexing.Lexer) *lexing.Token {
	r := x.Rune()
	if x.IsWhite(r) {
		panic("incorrect token start")
	}
	switch r {
	case '"':
		return lexing.LexString(x, tokString, r)
	case '`':
		return lexing.LexRawString(x, tokString)
	case ':':
		x.Next()
		return x.MakeToken(tokCollon)
	}

	if lexing.IsIdentLetter(r) {
		return lexing.LexIdent(x, tokIdent)
	}

	x.Errorf("illegal char %q", r)
	x.Next()
	return x.MakeToken(lexing.Illegal)
}

// Parse parse a line into an option.
func Parse(line string) (*Option, error) {
	r := strings.NewReader(line)
	x := lexing.MakeLexer("", r, lexLine)
	x.IsWhite = lexing.IsWhiteOrEndl
	toks := lexing.TokenAll(x)
	if errs := x.Errs(); errs != nil {
		return nil, fmt.Errorf("invalid %q: %s", line, errs[0])
	}

	if len(toks) != 4 {
		return nil, fmt.Errorf(
			"invalid %q: needs 4 tokens, got %d", line, len(toks),
		)
	}

	ident := toks[0]
	collon := toks[1]
	value := toks[2]
	eof := toks[3]

	if ident.Type != tokIdent {
		return nil, fmt.Errorf("want ident, got %s", tokStr(ident.Type))
	}
	if collon.Type != tokCollon {
		return nil, fmt.Errorf("want collon, got %s", tokStr(collon.Type))
	}
	if value.Type != tokString {
		return nil, fmt.Errorf("want string, got %s", tokStr(value.Type))
	}
	if eof.Type != lexing.EOF {
		return nil, fmt.Errorf("not end with EOF, but %s", tokStr(eof.Type))
	}

	v, err := strconv.Unquote(value.Lit)
	if err != nil {
		return nil, fmt.Errorf("invalid string %q: %s", value.Lit, err)
	}

	return &Option{
		Key:   ident.Lit,
		Value: v,
	}, nil

}
