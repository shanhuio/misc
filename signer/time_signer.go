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
	"encoding/binary"
	"time"
)

// TimeSigner signs the current time, or checks if a signed time
// is within a time window of the current time reading.
type TimeSigner struct {
	s      *Signer
	window time.Duration

	// TimeFunc is an optional function for reading teh current timestamp.
	// When it is nil, the TimeSinger uses time.Now()
	TimeFunc func() time.Time
}

func signTime(s *Signer, t time.Time) string {
	buf := make([]byte, timestampLen)
	binary.LittleEndian.PutUint64(buf, uint64(t.UnixNano()))
	return s.SignHex(buf)
}

// SignTime signes the current time.
func SignTime(key []byte) string {
	return signTime(New(key), time.Now())
}

// NewTimeSigner creates a new time singer.
func NewTimeSigner(key []byte, window time.Duration) *TimeSigner {
	if window < 0 {
		window = -window
	}
	return &TimeSigner{
		s:      New(key),
		window: window,
	}
}

// Token generates a signed token that has the current time reading.
func (s *TimeSigner) Token() string {
	return signTime(s.s, now(s.TimeFunc))
}

// Check checks if the timestamp is with in the time window.
func (s *TimeSigner) Check(token string) bool {
	ok, bs := s.s.CheckHex(token)
	if !ok {
		return false
	}
	if len(bs) != timestampLen {
		return false
	}

	t := time.Unix(0, int64(binary.LittleEndian.Uint64(bs)))
	timeNow := now(s.TimeFunc)
	return inWindow(t, timeNow, s.window)
}
