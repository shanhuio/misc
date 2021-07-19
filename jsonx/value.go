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

type value interface{}

type null struct {
	token *lexing.Token
}

type basic struct {
	lead  *lexing.Token
	token *lexing.Token
	value interface{}
}

type boolean struct {
	keyword *lexing.Token
}

type object struct {
	left    *lexing.Token
	entries []*objectEntry
	right   *lexing.Token
}

type objectKey struct {
	token *lexing.Token
	value interface{}
}

type objectEntry struct {
	key   *objectKey
	colon *lexing.Token
	value value
	comma *lexing.Token
}

type list struct {
	left    *lexing.Token
	entries []*listEntry
	right   *lexing.Token
}

type listEntry struct {
	value value
	comma *lexing.Token
}

type identList struct {
	entries []*lexing.Token
	dots    []*lexing.Token
}

type trunk struct {
	values []value
	semi   *lexing.Token
}
