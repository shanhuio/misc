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
	"fmt"
	"time"
)

// ForNow returns a human friendly string for describing a time point t
// relative to the current time.
func ForNow(t time.Time) string {
	return String(t, time.Now())
}

func timeDate(t time.Time) string {
	return t.Format("20060102")
}

// String returns a human friendly string for describing a time point t
// relative to the time point of now. When calculating time differences, it
// uses the time zone of the time point of now.
func String(t, now time.Time) string {
	t = t.In(now.Location())
	d := now.Sub(t)
	if d < 0 {
		return "in the future"
	}

	secs := int64(d / time.Second)
	if secs < 60 {
		return "just now"
	}

	mins := int64(d / time.Minute)
	if mins <= 1 {
		return "a minute ago"
	}
	if mins < 10 {
		return fmt.Sprintf("%d minutes ago", mins)
	}
	if mins < 60 {
		mins -= mins % 5
		return fmt.Sprintf("%d minutes ago", mins)
	}

	hours := int64(d / time.Hour)
	if hours <= 1 {
		return "an hour ago"
	}
	if hours <= 12 || timeDate(t) == timeDate(now) {
		return fmt.Sprintf("%d hours ago", hours)
	}

	days := int64(d / (time.Hour * 24))
	if days <= 1 {
		return "yesterday"
	}
	if days < 30 {
		return fmt.Sprintf("%d days ago", days)
	}

	months := int64(d / (time.Hour * 24 * 30))
	if months <= 1 {
		return "a month ago"
	}
	if months < 12 {
		return fmt.Sprintf("%d months ago", months)
	}

	years := int64(d / (time.Hour * 24 * 365))
	if years <= 1 {
		return "a year ago"
	}
	return fmt.Sprintf("%d years ago", years)
}
