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

package osutil

import (
	"testing"

	"io/ioutil"
	"os"
	"path"
)

func TestExist(t *testing.T) {
	ne := func(err error) {
		if err != nil {
			t.Fatal(err)
		}
	}

	d, err := ioutil.TempDir("", "osutil")
	ne(err)
	defer os.RemoveAll(d)

	ok, err := Exist(d)
	ne(err)
	if !ok {
		t.Errorf("dir %q should exist", d)
	}

	f := path.Join(d, "post")
	ne(ioutil.WriteFile(f, []byte("post"), 0600))

	ok, err = Exist(f)
	ne(err)
	if !ok {
		t.Errorf("file %q should exist", f)
	}

	ghost := path.Join(d, "ghost")
	ok, err = Exist(ghost)
	ne(err)
	if ok {
		t.Errorf("file %q should not exist", f)
	}
}
