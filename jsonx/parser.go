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
	"fmt"
	"io"

	"shanhu.io/text/lexing"
)

type parser struct {
	f string
	x lexing.Tokener
	*lexing.Parser
}

func newParser(f string, r io.Reader) (*parser, *lexing.Recorder) {
	p := &parser{f: f}
	x := tokener(f, r)
	rec := lexing.NewRecorder(x)
	p.x = lexing.NewCommentRemover(rec)
	p.Parser = lexing.NewParser(p.x, tokTypes)
	return p, rec
}

func (p *parser) seeOp(ops ...string) bool {
	t := p.Token()
	if t.Type != tokOperator {
		return false
	}
	for _, op := range ops {
		if t.Lit == op {
			return true
		}
	}
	return false
}

func tokenTypeStr(t *lexing.Token) string {
	if t.Type == tokOperator {
		return fmt.Sprintf("'%s'", t.Lit)
	}
	return typeStr(t.Type)
}

func (p *parser) expectOp(op string) *lexing.Token {
	if p.InError() {
		return nil
	}
	t := p.Token()
	if t.Type != tokOperator || t.Lit != op {
		p.CodeErrorfHere(
			"jsonx.expectOp", "expect '%s', got %s", op, tokenTypeStr(t),
		)
		return nil
	}
	return p.Shift()
}
