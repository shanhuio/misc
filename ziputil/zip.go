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
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// ZipFile creates a zip file of a single file.
func ZipFile(file string, w io.Writer) error {
	abs, err := filepath.Abs(file)
	if err != nil {
		return err
	}
	name := filepath.Base(abs)
	if name == "" {
		return fmt.Errorf("missing name for for the file")
	}

	info, err := os.Stat(file)
	if err != nil {
		return err
	}

	fin, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fin.Close()

	ar := zip.NewWriter(w)
	h := &zip.FileHeader{Name: name}
	h.SetMode(info.Mode())
	h.SetModTime(info.ModTime())

	zipFile, err := ar.CreateHeader(h)
	if err != nil {
		return err
	}

	if _, err := io.Copy(zipFile, fin); err != nil {
		return err
	}

	if err := fin.Close(); err != nil {
		return err
	}
	return ar.Close()
}

// ZipDir creates a zip file of a directory with all file in it.
func ZipDir(dir string, w io.Writer) error {
	ar := zip.NewWriter(w)
	walk := func(p string, info os.FileInfo, err error) error {
		rel, err := filepath.Rel(dir, p)
		if err != nil {
			return err
		}

		mod := info.Mode()
		t := info.ModTime()

		if info.IsDir() {
			h := &zip.FileHeader{Name: rel + "/"}
			h.SetMode(mod)
			h.SetModTime(t)
			_, err := ar.CreateHeader(h)
			return err
		}

		fin, err := os.Open(p)
		if err != nil {
			return err
		}
		defer fin.Close()

		h := &zip.FileHeader{Name: rel}
		h.SetMode(mod)
		h.SetModTime(t)

		w, err := ar.CreateHeader(h)
		if err != nil {
			return err
		}

		if _, err = io.Copy(w, fin); err != nil {
			return err
		}
		return fin.Close()
	}

	if err := filepath.Walk(dir, walk); err != nil {
		return err
	}

	return ar.Close()
}
