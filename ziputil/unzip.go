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
	"io"
	"os"
	"path/filepath"
)

// UnzipDir unzips a zip file into a directory.
// If the directory already exists, it removes all stuff in the directory
// if the clear flag is set to true.
func UnzipDir(dir string, r *zip.Reader, clear bool) error {
	if clear {
		err := os.RemoveAll(dir)
		if err != nil {
			return err
		}
	}

	for _, f := range r.File {
		mod := f.Mode()

		name := filepath.Join(dir, f.Name)
		if mod.IsDir() {
			if err := os.MkdirAll(name, mod); err != nil {
				return err
			}
			continue
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		if err := os.MkdirAll(filepath.Dir(name), 0700); err != nil {
			return err
		}

		fout, err := os.Create(name)
		if err != nil {
			return err
		}
		defer fout.Close()

		if err := fout.Chmod(mod); err != nil {
			return err
		}

		if _, err := io.Copy(fout, rc); err != nil {
			return err
		}

		if err := fout.Close(); err != nil {
			return err
		}
	}
	return nil
}
