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

package unixhttp

import (
	"net"
	"net/http"
	"os"

	"shanhu.io/misc/osutil"
)

// Listen creates a listener at the given unix domain socket address.
func Listen(p string) (net.Listener, error) {
	exist, err := osutil.IsSock(p)
	if err != nil {
		return nil, err
	}

	if exist {
		if err := os.Remove(p); err != nil {
			return nil, err
		}
	}
	lis, err := net.ListenUnix("unix", &net.UnixAddr{
		Name: p,
		Net:  "unix",
	})
	if err != nil {
		return nil, err
	}
	return lis, nil
}

// ListenAndServe listens and serves at the given unix domain socket
// path.
func ListenAndServe(p string, h http.Handler) error {
	lis, err := Listen(p)
	if err != nil {
		return err
	}
	return http.Serve(lis, h)
}
