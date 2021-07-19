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
	"fmt"

	"shanhu.io/misc/errcode"
)

type queryError struct {
	q   string
	err error
}

func (e *queryError) Error() string {
	return fmt.Sprintf("%s: query:\n%q", e.err, e.q)
}

// Error creates a query error if err is not nil.
// It returns nil if err is nil
func Error(q string, err error) error {
	if err == nil {
		return nil
	}
	qerr := &queryError{q: q, err: err}
	return errcode.Add(errcode.Internal, qerr)
}
