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

package testutil

import (
	"log"
)

// Failer is an interface to mark that a test has failed.
type Failer interface {
	Fail()
}

// PanicFailer fails with a log.Fatal call.
type PanicFailer struct{}

// Fail fails with a log.Fatal call.
func (pf *PanicFailer) Fail() {
	log.Fatal("test failed")
}
