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
	"io/ioutil"
	"os"
)

// File is a temp file.
type File struct {
	*os.File
	Name        string
	SkipCleanUp bool
}

// NewFile creates a new temp file.
func NewFile(dir, prefix string) (*File, error) {
	f, err := ioutil.TempFile(dir, prefix)
	if err != nil {
		return nil, err
	}
	return &File{
		File: f,
		Name: f.Name(),
	}, nil
}

// Reset is an alias for Seek(0, 0) on the file.
func (f *File) Reset() error {
	_, err := f.File.Seek(0, os.SEEK_SET)
	return err
}

// Remove removes the temp file.
func (f *File) Remove() error {
	return os.Remove(f.Name)
}

// Rename the file to another name.
func (f *File) Rename(p string) error {
	return os.Rename(f.Name, p)
}

// CleanUp closes the underlying temp file and removes it.
func (f *File) CleanUp() error {
	if f.SkipCleanUp {
		return nil
	}
	if err := f.File.Close(); err != nil {
		f.Remove() // remove anyways
		return err
	}
	return f.Remove()
}
