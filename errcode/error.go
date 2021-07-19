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

package errcode

import (
	"fmt"
)

// Error is a generic error with a string error code.
type Error struct {
	Code  string // code is the type of the error.
	error        // err is the error message, human friendly.
}

// Common general error codes
const (
	NotFound     = "not-found"
	InvalidArg   = "invalid-arg"
	Internal     = "internal"
	Unauthorized = "unauthorized"
	TimeOut      = "time-out"
)

// Add creates a new error with code as the error code.
func Add(code string, err error) *Error {
	return &Error{
		Code:  code,
		error: err,
	}
}

// Of returns the code of the error. For errors that
// do not have a code, it returns empty string.
func Of(err error) string {
	if codedErr, ok := err.(*Error); ok {
		return codedErr.Code
	}
	return ""
}

// IsNotFound checks if it is a not-found error.
func IsNotFound(err error) bool {
	return Of(err) == NotFound
}

// IsInvalidArg checks if it is an invalid argument error.
func IsInvalidArg(err error) bool {
	return Of(err) == InvalidArg
}

// IsInternal checks if it is an internal error.
func IsInternal(err error) bool {
	return Of(err) == Internal
}

// IsUnauthorized checks if it is an unauthorized error.
func IsUnauthorized(err error) bool {
	return Of(err) == Unauthorized
}

// IsTimeOut checks if it is a time-out error.
func IsTimeOut(err error) bool {
	return Of(err) == TimeOut
}

// Errorf creates an Error with the given error code.
func Errorf(code string, f string, args ...interface{}) *Error {
	return Add(code, fmt.Errorf(f, args...))
}

// AltErrorf replaces the message of e, but keeps the error code.
func AltErrorf(err error, msg string, args ...interface{}) error {
	cerr, ok := err.(*Error)
	if !ok {
		return fmt.Errorf(msg, args...)
	}
	return Errorf(cerr.Code, msg, args...)
}

// Annotate annotates an error but keeps the error code.
func Annotate(err error, msg string) error {
	return AltErrorf(err, "%s: %s", msg, err)
}

// Annotatef annotates an error with a formatted message but keeps the error
// code.
func Annotatef(err error, f string, args ...interface{}) error {
	return Annotate(err, fmt.Sprintf(f, args...))
}

// NotFoundf creates a new not-found error.
func NotFoundf(f string, args ...interface{}) *Error {
	return Errorf(NotFound, f, args...)
}

// InvalidArgf creates a new invalid arugment error.
func InvalidArgf(f string, args ...interface{}) *Error {
	return Errorf(InvalidArg, f, args...)
}

// Internalf creates a new internal error.
func Internalf(f string, args ...interface{}) *Error {
	return Errorf(Internal, f, args...)
}

// Unauthorizedf returns an error caused by an unauthrozied request.
func Unauthorizedf(f string, args ...interface{}) *Error {
	return Errorf(Unauthorized, f, args...)
}

// TimeOutf returns a new time-out error.
func TimeOutf(f string, args ...interface{}) *Error {
	return Errorf(TimeOut, f, args...)
}
