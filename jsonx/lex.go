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
	"io"

	"shanhu.io/text/lexing"
)

func lexOperator(x *lexing.Lexer, r rune) *lexing.Token {
	switch r {
	case '{', '}', '[', ']', ',', ':', '+', '-', '.':
		/* do nothing */
	case '/':
		r2 := x.Rune()
		if r2 == '/' || r2 == '*' {
			return lexing.LexComment(x)
		}
	case ';':
		return x.MakeToken(tokSemi)
	default:
		return nil
	}
	return x.MakeToken(tokOperator)
}

func lexJSONX(x *lexing.Lexer) *lexing.Token {
	r := x.Rune()
	if x.IsWhite(r) {
		panic("incorrect token start")
	}

	switch r {
	case '\n':
		x.Next()
		return x.MakeToken(tokEndl)
	case '"':
		return lexing.LexString(x, tokString, '"')
	case '`':
		return lexing.LexRawString(x, tokString)
	}

	if lexing.IsDigit(r) {
		return lexing.LexNumber(x, tokInt, tokFloat)
	}
	if lexing.IsIdentLetter(r) {
		return lexing.LexIdent(x, tokIdent)
	}

	x.Next()
	t := lexOperator(x, r)
	if t != nil {
		return t
	}

	x.CodeErrorf("jsonx.illegalChar", "illegal char %q", r)
	return x.MakeToken(lexing.Illegal)
}

var keywords = lexing.KeywordSet("true", "false", "null")

func tokener(f string, r io.Reader) lexing.Tokener {
	x := lexing.MakeLexer(f, r, lexJSONX)
	si := newSemiInserter(x)
	kw := lexing.NewKeyworder(si)
	kw.Ident = tokIdent
	kw.Keyword = tokKeyword
	kw.Keywords = keywords
	return kw
}
