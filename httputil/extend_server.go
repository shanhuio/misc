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
	"net/url"
	"strings"
)

// ExtendServer extends the server host string with an https:// prefix if it is
// not a localhost, or an http:// prefix if it is localhost. It only extends
// the server when the server string can be extends to be a valid URL.
func ExtendServer(s string) string {
	if strings.HasPrefix(s, "http://") || strings.HasPrefix(s, "https://") {
		return s
	}

	ret := "https://" + s
	u, err := url.Parse(ret)
	if err != nil {
		return s // not a valid URL
	}
	if u.Scheme != "https" {
		return s // not something that we think of
	}

	if u.Scheme == "https" {
		if u.Host == "localhost" || strings.HasPrefix(u.Host, "localhost:") {
			// testing on localhost, set the scheme to http
			u.Scheme = "http"
			return u.String()
		}
	}
	return ret
}
