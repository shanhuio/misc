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

package goload

import (
	"path"
	"strings"
)

type vendorLayer struct {
	p      string
	prefix string
	pkgs   map[string]string
}

func newVendorLayer(p string) *vendorLayer {
	return &vendorLayer{
		p:      p,
		prefix: path.Join(p, "vendor") + "/",
		pkgs:   make(map[string]string),
	}
}

func (ly *vendorLayer) addPkg(p string) {
	if p == ly.p {
		return
	}
	if !strings.HasPrefix(p, ly.prefix) {
		panic("not inside the vendor directory")
	}
	ly.pkgs[strings.TrimPrefix(p, ly.prefix)] = p
}

type vendorStack struct {
	layers []*vendorLayer
}

func (s *vendorStack) push(ly *vendorLayer) {
	s.layers = append(s.layers, ly)
}

func (s *vendorStack) pop() {
	n := len(s.layers)
	if n == 0 {
		panic("nothing to pop")
	}
	s.layers = s.layers[:n-1]
}

func (s *vendorStack) mapImport(p string) (string, bool) {
	n := len(s.layers)
	for i := n - 1; i >= 0; i-- {
		ly := s.layers[i]
		if mapped, found := ly.pkgs[p]; found {
			return mapped, true
		}
	}
	return p, false
}
