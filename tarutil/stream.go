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
	"io"
	"os"
	"time"

	"shanhu.io/misc/errcode"
)

// streamFile is a file (or a zip archive) to stream into a tar stream.
type streamFile struct {
	name    string // Name to write into the tar stream.
	file    string // File to read from file system.
	zip     bool   // If to read the file as a zip file.
	content []byte // Raw content; used only when File is empty string.

	meta Meta

	modTime time.Time
}

// Meta contains metadata of file
type Meta struct {
	Mode    int64
	UserID  int
	GroupID int
}

// ModeMeta creates a Meta with specific mode.
func ModeMeta(mode int64) *Meta { return &Meta{Mode: mode} }

func (f *streamFile) writeTo(tw *tar.Writer) error {
	if f.zip {
		return TarZipFile(tw, f.file)
	}

	if f.file != "" {
		file, err := os.Open(f.file)
		if err != nil {
			return err
		}
		defer file.Close()

		stat, err := file.Stat()
		if err != nil {
			return err
		}

		mode := f.meta.Mode
		if mode == 0 {
			mode = int64(stat.Mode()) & 0777
		}

		if err := tw.WriteHeader(&tar.Header{
			Name:    f.name,
			Size:    stat.Size(),
			Mode:    mode,
			Gid:     f.meta.GroupID,
			Uid:     f.meta.UserID,
			ModTime: f.modTime,
		}); err != nil {
			return err
		}
		_, err = io.Copy(tw, file)
		return err
	}

	if err := tw.WriteHeader(&tar.Header{
		Name:    f.name,
		Size:    int64(len(f.content)),
		Mode:    f.meta.Mode,
		Gid:     f.meta.GroupID,
		Uid:     f.meta.UserID,
		ModTime: f.modTime,
	}); err != nil {
		return err
	}
	if len(f.content) > 0 {
		if _, err := tw.Write(f.content); err != nil {
			return err
		}
	}
	return nil
}

// Stream is a tar stream of files (or zip files). Files are transfered in
// the order of adding.
type Stream struct {
	files   []*streamFile
	modTime time.Time
}

// NewStream create a new tar stream.
func NewStream() *Stream { return &Stream{modTime: time.Now()} }

// AddString adds a file of name into the stream,
// which content is str.
func (s *Stream) AddString(name string, m *Meta, str string) {
	s.AddBytes(name, m, []byte(str))
}

// AddBytes adds a file of name into the stream, which content is bs.
func (s *Stream) AddBytes(name string, m *Meta, bs []byte) {
	s.files = append(s.files, &streamFile{
		name:    name,
		content: bs,
		meta:    *m,
		modTime: s.modTime,
	})
}

// AddFile adds a file of name and mode into the stream,
// which content is read from file f.
func (s *Stream) AddFile(name string, m *Meta, f string) {
	s.files = append(s.files, &streamFile{
		name:    name,
		file:    f,
		meta:    *m,
		modTime: s.modTime,
	})
}

// AddZipFile adds a zip file into the stream.
func (s *Stream) AddZipFile(f string) {
	s.files = append(s.files, &streamFile{
		file: f,
		zip:  true,
	})
}

type countingWriter struct {
	w io.Writer
	n int64
}

func (w *countingWriter) Write(bs []byte) (int, error) {
	n, err := w.w.Write(bs)
	w.n += int64(n)
	return n, err
}

// WriteTo writes the entire stream out to w.
func (s *Stream) WriteTo(w io.Writer) (int64, error) {
	cw := &countingWriter{w: w}
	tw := tar.NewWriter(cw)
	for _, f := range s.files {
		if err := f.writeTo(tw); err != nil {
			return cw.n, errcode.Annotatef(err, "write %q", f.name)
		}
	}
	err := tw.Close() // Close() might flush stuff and update cw.n
	return cw.n, err
}
