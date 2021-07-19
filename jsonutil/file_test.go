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

package jsonutil

import (
	"testing"

	"io/ioutil"
	"os"
	"reflect"
)

type testSturct struct {
	Number  int
	Boolean bool
	Text    string
}

var testWriteData = &testSturct{
	Text: "be stronger",
}

const testWriteReadable = `{
  "Number": 0,
  "Boolean": false,
  "Text": "be stronger"
}
`

func TestReadFile_notExist(t *testing.T) {
	const filename = "testdata/rumpelstilzchen"
	obj := &struct{}{}
	if err := ReadFile(filename, obj); err == nil {
		t.Errorf(
			"Read %s: want not-exist err, got nil",
			filename,
		)
	} else if !os.IsNotExist(err) {
		t.Errorf(
			"Read %s: want not-exist err, got %s",
			filename, err,
		)
	}
}

func TestReadFile_notJson(t *testing.T) {
	const filename = "testdata/invalid.json"
	obj := &struct{}{}
	if err := ReadFile(filename, obj); err == nil {
		t.Errorf(
			"Read %s: want unmarshal error, got %s",
			filename, err,
		)
	}
}

func TestReadFile(t *testing.T) {
	data := new(testSturct)
	const filename = "testdata/stronger.json"
	if err := ReadFile(filename, data); err != nil {
		t.Fatalf(
			"Read %q: got error: %s", filename, err,
		)
	}

	if !reflect.DeepEqual(data, testWriteData) {
		t.Errorf(
			"Read %s: want %v, got %v",
			filename, testWriteData, data,
		)
	}
}

func TestWriteFile(t *testing.T) {
	f, err := ioutil.TempFile("", "jsonfile-test")
	if err != nil {
		t.Fatal(err)
	}
	filename := f.Name()
	f.Close()
	defer os.Remove(filename)

	if err := WriteFile(filename, testWriteData); err != nil {
		t.Fatalf("Failed to Write %s: %s", filename, err)
	}
	dat := new(testSturct)
	if err := ReadFile(filename, dat); err != nil {
		t.Fatalf("Failed to Read %s: %s", filename, err)
	}

	if !reflect.DeepEqual(dat, testWriteData) {
		t.Errorf("expect %v, got %v", testWriteData, dat)
	}
}

func TestWriteFileReadable(t *testing.T) {
	f, err := ioutil.TempFile("", "jsonfile-test")
	if err != nil {
		t.Fatal(err)
	}
	filename := f.Name()
	f.Close()
	defer os.Remove(filename)

	if err := WriteFileReadable(filename, testWriteData); err != nil {
		t.Fatalf("Failed to WriteReadable %s: %s", filename, err)
	}

	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	got := string(bs)
	if got != testWriteReadable {
		t.Errorf("WriteReadable want %q, got %q", testWriteReadable, got)
	}
}
