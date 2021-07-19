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
	"strconv"

	"shanhu.io/text/lexing"
)

func parseStringValue(p *parser, t *lexing.Token) string {
	v, err := strconv.Unquote(t.Lit)
	if err != nil {
		p.CodeErrorf(
			t.Pos, "jsonx.stringLit", "invalid string: %s", err.Error(),
		)
		return ""
	}
	return v
}

func parseFloatValue(p *parser, t *lexing.Token) float64 {
	v, err := strconv.ParseFloat(t.Lit, 64)
	if err != nil {
		p.CodeErrorf(
			t.Pos, "jsonx.floatLit", "invalid float: %s", err.Error(),
		)
		return 0
	}
	return v
}

func parseObjectEntries(p *parser) []*objectEntry {
	var entries []*objectEntry
	for !p.seeOp("}") {
		if !(p.See(tokIdent) || p.See(tokString)) {
			p.CodeErrorfHere("jsonx.expectObjectEntry", "expect object entry")
			break
		}

		key := &objectKey{token: p.Shift()}
		if key.token.Type == tokString {
			key.value = parseStringValue(p, key.token)
		}
		colon := p.expectOp(":")
		v := parseValue(p)
		entry := &objectEntry{
			key:   key,
			colon: colon,
			value: v,
		}

		if p.seeOp(",") {
			entry.comma = p.Shift()
		} else if !p.seeOp("}") {
			p.expectOp(",")
		}
		entries = append(entries, entry)

		if p.InError() {
			break
		}
	}

	return entries
}

func parseListEntries(p *parser) []*listEntry {
	var entries []*listEntry
	for !p.seeOp("]") {
		v := parseValue(p)
		entry := &listEntry{value: v}
		if p.seeOp(",") {
			entry.comma = p.Shift()
		} else if !p.seeOp("]") {
			p.expectOp(",")
		}
		entries = append(entries, entry)
		if p.InError() {
			break
		}
	}

	return entries
}

func parseIdentList(p *parser) *identList {
	lst := new(identList)
	for {
		tok := p.Expect(tokIdent)
		if tok == nil {
			return lst
		}
		lst.entries = append(lst.entries, tok)
		if !p.seeOp(".") {
			break
		}
		lst.dots = append(lst.dots, p.Shift())
	}
	return lst
}

func parseValue(p *parser) value {
	switch {
	case p.See(tokKeyword):
		kw := p.Shift()
		if kw.Lit == "true" || kw.Lit == "false" {
			return &boolean{keyword: kw}
		}
		if kw.Lit == "null" {
			return &null{token: kw}
		}
		p.CodeErrorf(
			kw.Pos, "jsonx.unexpectedKeyword",
			"unexpected keyword '%s'", kw.Lit,
		)
		return nil
	case p.See(tokString):
		tok := p.Shift()
		return &basic{
			token: tok,
			value: parseStringValue(p, tok),
		}
	case p.See(tokInt):
		return &basic{token: p.Shift()}
	case p.See(tokFloat):
		tok := p.Shift()
		return &basic{
			token: tok,
			value: parseFloatValue(p, tok),
		}
	case p.seeOp("+", "-"):
		lead := p.Shift()
		if p.See(tokInt) || p.See(tokFloat) {
			return &basic{
				lead:  lead,
				token: p.Shift(),
			}
		}
		t := p.Token()
		p.CodeErrorf(
			t.Pos, "jsonx.expectNumber",
			"expect number, got %s", tokenTypeStr(t),
		)
		return nil
	case p.seeOp("{"):
		left := p.Shift()
		entries := parseObjectEntries(p)
		right := p.expectOp("}")
		return &object{
			left:    left,
			entries: entries,
			right:   right,
		}
	case p.seeOp("["):
		left := p.Shift()
		entries := parseListEntries(p)
		right := p.expectOp("]")
		return &list{
			left:    left,
			entries: entries,
			right:   right,
		}
	case p.See(tokIdent):
		return parseIdentList(p)
	default:
		t := p.Token()
		p.CodeErrorf(
			t.Pos, "jsonx.expectOperand",
			"expect an operand, got %s", typeStr(t.Type),
		)
		return nil
	}
}

func parseTrunk(p *parser) *trunk {
	t := &trunk{}
	for !p.See(tokSemi) {
		v := parseValue(p)
		t.values = append(t.values, v)
	}
	t.semi = p.Expect(tokSemi)
	return t
}
