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

package osutil

import (
	"io/ioutil"
	"os"
	"strings"
)

// ReadTokenFile reads a token string from a file.
func ReadTokenFile(f string) (string, error) {
	bs, err := ioutil.ReadFile(f)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(bs)), nil
}

// ReadOptionalTokenFile reads an optional token file.
func ReadOptionalTokenFile(f string) (string, bool, error) {
	ret, err := ReadTokenFile(f)
	if err != nil {
		if os.IsNotExist(err) {
			return "", false, nil
		}
		return "", false, err
	}
	return ret, true, nil
}
