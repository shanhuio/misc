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

package pathutil

import (
	"fmt"
)

type codeError struct {
	code string
	err  error
}

func (err *codeError) Error() string {
	return err.err.Error()
}

func codeErrorf(code string, f string, args ...interface{}) error {
	err := fmt.Errorf(f, args...)
	return &codeError{
		code: code,
		err:  err,
	}
}

func isCodeError(err error, code string) bool {
	codeErr, ok := err.(*codeError)
	if !ok {
		return false
	}

	return codeErr.code == code
}

func errCode(err error) string {
	codeErr, ok := err.(*codeError)
	if !ok {
		return ""
	}
	return codeErr.code
}

// IsNotExist returns true if the error means that a path does not exist as a
// file. It could be that the file is a directory/tree.
func IsNotExist(err error) bool {
	if err == nil {
		return false
	}
	code := errCode(err)
	return code == "not-exists" || code == "is-dir"
}

// NotExist creates an error where the given path does not exist.
func NotExist(name string) error {
	return codeErrorf("not-exists", "%q not found", name)
}

// Exist creates an error where the given path already exists.
func Exist(name string) error {
	return codeErrorf("exists", "%q exists", name)
}

// NotAbs creates an error where the given path is not absolute.
func NotAbs(name string) error {
	return codeErrorf("not-abs", "%q is not an absolute path", name)
}

// NotDir creates an error where the given path is not a directory.
func NotDir(name string) error {
	return codeErrorf("not-dir", "%q is not a directory", name)
}

// IsDir creates an error where the given path is a directory.
func IsDir(name string) error {
	return codeErrorf("is-dir", "%q is a directory", name)
}

// Invalid creates an error where the given path is invalid.
func Invalid(name string) error {
	return codeErrorf("invalid", "%q is an invalid path", name)
}

// ReadOnly creates an error where the given path is read-only.
func ReadOnly(name string) error {
	return codeErrorf("read-only", "%q is readonly", name)
}
