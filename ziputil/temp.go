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

	"shanhu.io/misc/tempfile"
)

// OpenInTemp copies all bytes of r into a temp file, and then opens the
// creates a zip reader that is backed by f.
func OpenInTemp(r io.Reader, tmp *tempfile.File) (*zip.Reader, error) {
	n, err := io.Copy(tmp, r)
	if err != nil {
		return nil, err
	}
	if err := tmp.Reset(); err != nil {
		return nil, err
	}
	return zip.NewReader(tmp, n)
}
