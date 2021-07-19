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

// Row is a result with the row and the query.
type Row struct {
	Query string
	*sql.Row
}

// Scan scans a row into values.
func (r *Row) Scan(dest ...interface{}) (bool, error) {
	err := r.Row.Scan(dest...)
	if err == nil {
		return true, nil
	}
	if err == sql.ErrNoRows {
		return false, nil
	}
	return false, Error(r.Query, err)
}

// Rows is a result with the rows and the query.
type Rows struct {
	Query string
	*sql.Rows
}

// Close closes the rows result.
func (r *Rows) Close() error {
	return Error(r.Query, r.Rows.Close())
}

// Err returns the error.
func (r *Rows) Err() error {
	return Error(r.Query, r.Rows.Err())
}

// Scan scans a row into values.
func (r *Rows) Scan(dest ...interface{}) error {
	return Error(r.Query, r.Rows.Scan(dest...))
}
