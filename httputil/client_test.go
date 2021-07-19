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
	"testing"

	"io/ioutil"
)

func TestClientGetCode(t *testing.T) {
	s := newHelloServer()
	c, err := NewClient(s.URL)
	if err != nil {
		t.Fatal(err)
	}
	got, err := c.GetCode("/")
	if err != nil {
		t.Fatal(err)
	}
	if got != 200 {
		t.Errorf("want 200, got %d", got)
	}

	got, err = c.GetCode("/secret")
	if err != nil {
		t.Fatal(err)
	}
	if got != 403 {
		t.Errorf("want 403, got %d", got)
	}
}

func TestClientGet(t *testing.T) {
	s := newHelloServer()
	c, err := NewClient(s.URL)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := c.Get("/")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	got, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != testHelloMessage {
		t.Errorf("got %q, want %q", string(got), testHelloMessage)
	}
}
