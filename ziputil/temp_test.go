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

package ziputil

import (
	"testing"

	"bytes"
	"io/ioutil"

	"shanhu.io/misc/tempfile"
)

func TestOpenInTemp(t *testing.T) {
	ne := func(err error) {
		if err != nil {
			t.Fatal(err)
		}
	}

	bs, err := ioutil.ReadFile("testdata/testfile.zip")
	ne(err)

	f, err := tempfile.NewFile("", "ziputil")
	ne(err)
	defer f.CleanUp()

	r, err := OpenInTemp(bytes.NewReader(bs), f)
	ne(err)

	if len(r.File) != 1 {
		t.Fatal("want 1 file in testfile.zip")
	}

	got := r.File[0].Name
	want := "testfile"
	if got != want {
		t.Errorf("file name want %q, got %q", want, got)
	}
}
