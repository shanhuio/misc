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

package trie

// Trie is a trie. Each branch split is a string rather than a letter.
type Trie struct {
	root *node
}

// New creates a new trie.
func New() *Trie {
	return &Trie{
		root: newNode(),
	}
}

// Add adds a new routed value into the trie.
func (t *Trie) Add(route []string, value string) bool {
	if value == "" {
		panic("value cannot be empty")
	}
	return t.root.add(route, value)
}

// FindExact looks for the value of a particular route.
// Returns empty string if not found.
func (t *Trie) FindExact(route []string) string {
	n, v := t.root.find(route)
	if n != len(route) {
		return ""
	}
	return v
}

// Find looks for the longest prefix match for the route.
func (t *Trie) Find(route []string) (match []string, v string) {
	n, v := t.root.find(route)
	if v == "" {
		return nil, ""
	}
	return route[:n], v
}
