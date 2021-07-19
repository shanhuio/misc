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

package tarutil

import (
	"archive/tar"
	"archive/zip"
	"io"

	"shanhu.io/misc/errcode"
)

func copyZipFile(w io.Writer, f *zip.File) error {
	rc, err := f.Open()
	if err != nil {
		return err
	}
	if _, err := io.Copy(w, rc); err != nil {
		rc.Close()
		return err
	}
	return rc.Close()
}

// TarZipFile puts all files from a zip file into a tar stream.
func TarZipFile(tw *tar.Writer, p string) error {
	z, err := zip.OpenReader(p)
	if err != nil {
		return errcode.Annotate(err, "open zip file")
	}

	for _, f := range z.File {
		stat := f.FileInfo()
		tarStat, err := tar.FileInfoHeader(stat, "")
		if err != nil {
			return errcode.Annotatef(err, "tar stat for: %q", f.Name)
		}
		tarStat.Name = f.Name
		if err := tw.WriteHeader(tarStat); err != nil {
			return errcode.Annotatef(err, "write header: %q", f.Name)
		}
		if err := copyZipFile(tw, f); err != nil {
			return errcode.Annotatef(err, "copy zip file: %q", f.Name)
		}
	}

	return nil
}
