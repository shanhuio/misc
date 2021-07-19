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
	"shanhu.io/text/lexing"
)

const (
	tokKeyword = iota
	tokIdent
	tokString
	tokInt
	tokFloat
	tokOperator
	tokSemi
	tokEndl
)

var tokTypes = func() *lexing.Types {
	t := lexing.NewTypes()
	for _, e := range []struct {
		t    int
		name string
	}{
		{tokKeyword, "keyword"},
		{tokIdent, "identifier"},
		{tokString, "string"},
		{tokInt, "integer"},
		{tokFloat, "float"},
		{tokOperator, "operator"},
		{tokEndl, "end-line"},
		{tokSemi, "end-line"},
	} {
		t.Register(e.t, e.name)
	}
	return t
}()

func typeStr(t int) string { return tokTypes.Name(t) }
