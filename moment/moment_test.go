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

package moment

import (
	"testing"

	"time"
)

func TestString(t *testing.T) {
	now := time.Now()
	for _, test := range []struct {
		dur  time.Duration
		want string
	}{
		{0, "just now"},
		{time.Second, "just now"},
		{3 * time.Second, "just now"},
		{59 * time.Second, "just now"},
		{time.Minute, "a minute ago"},
		{1*time.Minute + 30*time.Second, "a minute ago"},
		{3 * time.Minute, "3 minutes ago"},
		{5*time.Minute + 27*time.Second, "5 minutes ago"},
		{12 * time.Minute, "10 minutes ago"},
		{59 * time.Minute, "55 minutes ago"},
		{time.Hour, "an hour ago"},
		{time.Hour + 59*time.Minute, "an hour ago"},
		{time.Hour * 2, "2 hours ago"},
		{time.Hour * 24, "yesterday"},
		{time.Hour * 47, "yesterday"}, // this might not be so correct
	} {
		got := String(now.Add(-test.dur), now)
		if got != test.want {
			t.Errorf(
				"moment string for %s: got %q, want %q",
				test.dur, got, test.want,
			)
		}
	}

	for _, test := range []struct {
		hour int
		want string
	}{
		{23, "13 hours ago"},
		{14, "13 hours ago"},
		{1, "yesterday"},
		{12, "yesterday"},
	} {
		now := time.Date(1, 1, 1, test.hour, 0, 0, 0, time.UTC)
		dur := 13 * time.Hour
		got := String(now.Add(-dur), now)
		if got != test.want {
			t.Errorf(
				"moment string for %s afer %q: got %q, want %q",
				dur, now, got, test.want,
			)
		}
	}

	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Fatal(err)
	}
	for _, test := range []struct {
		hour int
		want string
	}{
		{23, "13 hours ago"},
		{14, "13 hours ago"},
		{1, "yesterday"},
		{12, "yesterday"},
	} {
		utcTime := time.Date(1, 1, 1, test.hour, 0, 0, 0, time.UTC)
		dur := 13 * time.Hour
		nyTime := utcTime.Add(-dur).In(loc)
		got := String(nyTime, utcTime)
		if got != test.want {
			t.Errorf(
				"moment string for %s afer %q, %q: got %q, want %q",
				dur, utcTime, nyTime.In(time.UTC), got, test.want,
			)
		}
	}
}
