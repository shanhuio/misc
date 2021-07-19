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

func TestStates(t *testing.T) {
	s := NewSessions(nil, time.Second)
	state := s.NewState()
	if !s.CheckState(state) {
		t.Errorf("check on state %q failed", state)
	}

	if s.CheckState("") {
		t.Errorf("check on empty state is passing")
	}
}

func TestStatesExpire(t *testing.T) {
	const ttl = time.Second
	s := NewSessions(nil, ttl)
	now := time.Unix(0, 0)
	s.TimeFunc = func() time.Time { return now }
	state := s.NewState()
	t.Log("state: ", state)

	now = now.Add(ttl).Add(-time.Nanosecond)
	if !s.CheckState(state) {
		t.Errorf("check on state %q failed", state)
	}

	now = now.Add(time.Nanosecond)
	if s.CheckState(state) {
		t.Errorf("check passed, should fail because of time out")
	}
}
