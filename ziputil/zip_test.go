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

	"archive/zip"
	"io"
	"io/ioutil"
	"os"
	"path"
	"reflect"

	"shanhu.io/misc/osutil"
	"shanhu.io/misc/tempfile"
)

func testDiffFile(t *testing.T, f1, f2 string) bool {
	ne := func(err error) {
		if err != nil {
			t.Fatal(err)
		}
	}

	bs1, err := ioutil.ReadFile(f1)
	ne(err)

	bs2, err := ioutil.ReadFile(f2)
	ne(err)

	if !reflect.DeepEqual(bs1, bs2) {
		return false
	}

	s1, err := os.Stat(f1)
	ne(err)

	s2, err := os.Stat(f2)
	ne(err)

	if s1.Mode() != s2.Mode() {
		return false
	}
	return true
}

func TestZipFile(t *testing.T) {
	ne := func(err error) {
		if err != nil {
			t.Fatal(err)
		}
	}

	temp, err := tempfile.NewFile("", "ziputil")
	ne(err)
	defer temp.CleanUp()

	const p = "testdata/testfile"
	ne(ZipFile(p, temp))

	size, err := temp.Seek(0, io.SeekCurrent)
	ne(err)

	ne(temp.Reset())

	output, err := ioutil.TempDir("", "ziputil")
	ne(err)

	defer os.RemoveAll(output)

	z, err := zip.NewReader(temp, size)
	ne(err)

	ne(UnzipDir(output, z, true))

	outPath := path.Join(output, "testfile")
	if !testDiffFile(t, outPath, p) {
		t.Error("zip loop back failed")
	}
}

func TestZipDir(t *testing.T) {
	ne := func(err error) {
		if err != nil {
			t.Fatal(err)
		}
	}

	temp, err := tempfile.NewFile("", "ziputil")
	ne(err)
	defer temp.CleanUp()

	const p = "testdata/testdir"
	ne(ZipDir(p, temp))

	size, err := temp.Seek(0, io.SeekCurrent)
	ne(err)

	ne(temp.Reset())

	output, err := ioutil.TempDir("", "ziputil")
	ne(err)
	defer os.RemoveAll(output)

	z, err := zip.NewReader(temp, size)
	ne(err)

	ne(UnzipDir(output, z, true))

	for _, name := range []string{
		"bin-file", "private-file", "text-file",
	} {
		outPath := path.Join(output, name)
		target := path.Join(p, name)
		if !testDiffFile(t, outPath, target) {
			t.Errorf("zip loop back failed for file %q", name)
		}
	}
}

func testClearDir(t *testing.T, clear bool) {
	ne := func(err error) {
		if err != nil {
			t.Fatal(err)
		}
	}

	z, err := zip.OpenReader("testdata/testfile.zip")
	ne(err)
	defer z.Close()

	output, err := ioutil.TempDir("", "ziputil")
	ne(err)
	defer os.RemoveAll(output)

	ob := path.Join(output, "native-file")
	ne(ioutil.WriteFile(ob, []byte("lived here long time ago"), 0600))

	ne(UnzipDir(output, &z.Reader, clear))

	exist, err := osutil.Exist(ob)
	ne(err)

	if clear && exist {
		t.Error("should clear directory, but still see the file")
	}
	if !clear && !exist {
		t.Error("should preserve the file, but lost")
	}
}

func TestClearDir(t *testing.T) { testClearDir(t, true) }

func TestNoClearDir(t *testing.T) { testClearDir(t, false) }
