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

package netutil

import (
	"net"
	"time"
)

type keepAliveListner struct {
	*net.TCPListener
}

func (ln keepAliveListner) Accept() (c net.Conn, err error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return
	}

	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(3 * time.Minute)
	return tc, nil
}

// WrapKeepAlive wraps the listener. If the listener is a TCP listener, it
// sets keep alive to 3 minute.
func WrapKeepAlive(ln net.Listener) net.Listener {
	tcpLis, ok := ln.(*net.TCPListener)
	if !ok {
		return ln
	}
	return keepAliveListner{tcpLis}
}
