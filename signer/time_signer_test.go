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
	"testing"

	"time"
)

func TestTimeSigner(t *testing.T) {
	s := NewTimeSigner(nil, time.Second*5)
	now := time.Unix(0, 0)
	s.TimeFunc = func() time.Time { return now }

	token := s.Token()
	if !s.Check(token) {
		t.Errorf("token should be valid")
	}

	now = time.Unix(1, 0)
	if !s.Check(token) {
		t.Errorf("token should be still valid")
	}

	now = time.Unix(10, 0)
	if s.Check(token) {
		t.Errorf("token should be invalid")
	}

	now = time.Unix(-10, 0)
	if s.Check(token) {
		t.Errorf("token should be invalid")
	}
}
