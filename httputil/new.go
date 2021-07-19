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
)

// NewClientMust creates a client and panics on error.
func NewClientMust(s string) *Client { return NewTokenClientMust(s, "") }

// NewTokenClientMust creates a client with auth token and panics on error.
func NewTokenClientMust(s, tok string) *Client {
	c, err := NewTokenClient(s, tok)
	if err != nil {
		panic(err)
	}
	return c
}

// NewClient creates a new client.
func NewClient(s string) (*Client, error) {
	u, err := url.Parse(s)
	if err != nil {
		return nil, err
	}
	return &Client{Server: u}, nil
}

// NewTokenClient creates a new client with a Bearer token.
func NewTokenClient(s, tok string) (*Client, error) {
	c, err := NewClient(s)
	if err != nil {
		return nil, err
	}
	c.Token = tok
	return c, nil
}

// NewUnixClient creates a new client that always goes to a particular
// unix domain socket.
func NewUnixClient(sockAddr string) *Client {
	return &Client{
		Server:    &url.URL{Scheme: "http", Host: "unix.sock"},
		Transport: unixSockTransport(sockAddr),
	}
}
