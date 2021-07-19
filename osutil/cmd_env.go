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
	"fmt"
	"os"
	"os/exec"
)

// CmdAddEnv adds an environment variable to cmd and returns true. If v is an
// empty string, nothing is added, and it returns false.
func CmdAddEnv(cmd *exec.Cmd, k, v string) bool {
	if v == "" {
		return false
	}
	cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v))
	return true
}

// CmdCopyEnv copies the value of environment variable k to cmd. If the value
// is empty, returns false; otherwise returns true.
func CmdCopyEnv(cmd *exec.Cmd, k string) bool {
	return CmdAddEnv(cmd, k, os.Getenv(k))
}
