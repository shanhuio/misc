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

package tempfile

import (
	"testing"

	"io/ioutil"
	"os"
	"reflect"

	"shanhu.io/misc/osutil"
)

func testFileExist(t *testing.T, name string) bool {
	ret, err := osutil.Exist(name)
	if err != nil {
		t.Fatal(err)
		return false
	}
	return ret
}

func TestFileReadBack(t *testing.T) {
	ne := func(err error) {
		if err != nil {
			t.Fatal(err)
		}
	}
	f, err := NewFile("", "tempfile")
	ne(err)
	defer f.CleanUp()

	msg := []byte("hello")

	_, err = f.Write(msg)
	ne(err)

	ne(f.Reset())

	readBack, err := ioutil.ReadAll(f)
	ne(err)

	if !reflect.DeepEqual(msg, readBack) {
		t.Errorf("want %q, got %q", string(msg), string(readBack))
	}
}

func TestFileCleanUp(t *testing.T) {
	ne := func(err error) {
		if err != nil {
			t.Fatal(err)
		}
	}

	testFunc := func() string {
		f, err := NewFile("", "tempfile")
		ne(err)
		defer f.CleanUp()

		return f.Name
	}

	f := testFunc()
	if testFileExist(t, f) {
		t.Errorf("test file %q still exist", f)
	}
}

func TestFileRename(t *testing.T) {
	ne := func(err error) {
		if err != nil {
			t.Fatal(err)
		}
	}

	testFunc := func() (string, string) {
		f, err := NewFile("", "tempfile")
		ne(err)
		defer f.CleanUp()

		f.SkipCleanUp = true

		target, err := NewFile("", "tempfile")
		ne(err)
		defer target.CleanUp()

		ne(f.Rename(target.Name))

		return f.Name, target.Name
	}

	f1, f2 := testFunc()
	if testFileExist(t, f1) {
		t.Errorf("test file1 %q still exist", f1)
	}
	if testFileExist(t, f2) {
		t.Errorf("test file2 %q still exist", f2)
	}
}

func TestFileSkipCleanUp(t *testing.T) {
	ne := func(err error) {
		if err != nil {
			t.Fatal(err)
		}
	}

	testFunc := func() string {
		f, err := NewFile("", "tempfile")
		ne(err)
		f.SkipCleanUp = true
		defer f.CleanUp()

		return f.Name
	}

	f := testFunc()
	if !testFileExist(t, f) {
		t.Errorf("test file %q should exist", f)
	}

	ne(os.Remove(f))
}
