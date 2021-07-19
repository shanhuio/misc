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

package errcode

import (
	"os"
)

// FromOS converts os package errors into errcode errors.
func FromOS(err error) error {
	if err == nil {
		return err
	}
	if os.IsNotExist(err) {
		return Add(NotFound, err)
	}
	if os.IsPermission(err) {
		return Add(Unauthorized, err)
	}
	if os.IsTimeout(err) {
		return Add(TimeOut, err)
	}
	return err
}
