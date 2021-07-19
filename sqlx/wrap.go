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

package sqlx

import (
	"database/sql"
)

type wrap struct {
	conn
}

// X executes a query string.
func (w *wrap) X(q string, args ...interface{}) (sql.Result, error) {
	res, err := w.conn.Exec(q, args...)
	return res, Error(q, err)
}

// Q1 executes a query string that expects one single row as the return.
func (w *wrap) Q1(q string, args ...interface{}) *Row {
	return &Row{
		Query: q,
		Row:   w.conn.QueryRow(q, args...),
	}
}

// Q queries the database with a query string.
func (w *wrap) Q(q string, args ...interface{}) (*sql.Rows, error) {
	res, err := w.conn.Query(q, args...)
	if err != nil {
		return nil, Error(q, err)
	}
	return res, nil
}
