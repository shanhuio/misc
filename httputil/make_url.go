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
	"path"
)

func makeURL(base *url.URL, p string) (string, error) {
	u := *base
	up, err := url.Parse(p)
	if err != nil {
		return "", err
	}

	// append two paths
	u.Path = path.Join(u.Path, up.Path)
	u.RawQuery = up.RawQuery
	u.Fragment = up.Fragment
	return u.String(), nil
}
