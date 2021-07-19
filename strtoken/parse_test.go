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

package strtoken

import (
	"testing"
)

func TestParseLine(t *testing.T) {
	o := func(line string, args ...string) {
		strs, errs := Parse(line)
		if errs != nil {
			t.Errorf("Parse(%q): unexpected errors", line)
			for _, err := range errs {
				t.Log(err)
			}
			return
		}

		if len(strs) != len(args) {
			t.Errorf("Parse(%q): expect %d args, got %d",
				line, len(args), len(strs),
			)
			return
		}

		for i, s := range strs {
			if s != args[i] {
				t.Errorf("Parse(%q), arg %d: expect %q, got %q",
					line, i, args[i], s,
				)
			}
		}
	}
	o("")
	o("a", "a")
	o(`"a"`, "a")
	o(`/something`, "/something")
	o(`ls /x_file`, "ls", "/x_file")
	o("       ls \t\t a", "ls", "a")
	o(`"a-b" something`, "a-b", "something")
	o(`a-b`, "a-b")
	o(`?`, "?")
	o(`!x`, "!x")
	o(`mkdir -p /root/.ssh`, "mkdir", "-p", "/root/.ssh")

	e := func(line string) {
		_, errs := Parse(line)
		if errs == nil {
			t.Errorf("Parse(%q): expect error but passed", line)
		}
	}

	e(`"`)
	e(`"asdf" asdf "xx`)
}
