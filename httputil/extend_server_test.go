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

package httputil

import (
	"testing"
)

func TestExtendServer(t *testing.T) {
	o := func(s, want string) {
		got := ExtendServer(s)
		if got != want {
			t.Errorf("extendServer(%q), got %q want %q", s, got, want)
		}
	}

	o("http://localhost", "http://localhost")
	o("https://shanhu.io", "https://shanhu.io")
	o("localhost", "http://localhost")
	o("localhost:3356", "http://localhost:3356")
	o("shanhu.io", "https://shanhu.io")
}
