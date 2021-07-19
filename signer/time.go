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

package signer

import (
	"time"
)

const timestampLen = 8

func now(f func() time.Time) time.Time {
	if f == nil {
		return time.Now()
	}
	return f()
}

func inWindow(t, tnow time.Time, w time.Duration) bool {
	tstart := tnow.Add(-w)
	tend := tnow.Add(w)
	return t.After(tstart) && t.Before(tend)
}
