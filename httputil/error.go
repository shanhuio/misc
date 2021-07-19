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

package httputil

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"shanhu.io/misc/errcode"
)

func isSuccess(resp *http.Response) bool {
	return resp.StatusCode/100 == 2
}

type httpError struct {
	StatusCode int
	Status     string
	Body       string
}

func (err *httpError) Error() string {
	if err.Body != "" {
		return fmt.Sprintf("%s - %s", err.Status, err.Body)
	}
	return err.Status
}

// ErrorStatusCode returns the status code is it is an HTTP error.
func ErrorStatusCode(err error) int {
	herr, ok := err.(*httpError)
	if !ok {
		return 0
	}
	return herr.StatusCode
}

// AddErrCode adds error code to an error given the http status.
func AddErrCode(statusCode int, err error) error {
	switch statusCode {
	case http.StatusNotFound:
		err = errcode.Add(errcode.NotFound, err)
	case http.StatusUnauthorized, http.StatusForbidden:
		err = errcode.Add(errcode.Unauthorized, err)
	case http.StatusBadRequest:
		err = errcode.Add(errcode.InvalidArg, err)
	}
	return err
}

// RespError returns the error from an HTTP response.
func RespError(resp *http.Response) error {
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	herr := &httpError{
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Body:       strings.TrimSpace(string(bs)),
	}
	return AddErrCode(resp.StatusCode, herr)
}
