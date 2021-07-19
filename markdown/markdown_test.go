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

package markdown

import (
	"testing"
)

func TestToHTML(t *testing.T) {
	for _, d := range []struct {
		in, out string
	}{
		{"", ""},
	} {
		out := ToHTML([]byte(d.in))
		if string(out) != d.out {
			t.Errorf("with title for %q", d.in)
			t.Logf("got output %q", string(out))
			t.Logf("want output %q", string(d.out))
		}
	}

	for _, d := range []struct {
		in, out, title string
	}{
		{"", "", ""},
		{"# Hello", "", "Hello"},
		{"# Hello\nContent.\n", "<p>Content.</p>\n", "Hello"},
	} {
		title, out := ToHTMLWithTitle([]byte(d.in))
		if string(out) != d.out || title != d.title {
			t.Errorf("with title for %q", d.in)
			t.Logf("got title %q and output %q", title, string(out))
			t.Logf("want title %q and output %q", d.title, string(d.out))
		}
	}
}
