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

type node struct {
	subs  map[string]*node
	value string // empty if not a leaf node
}

func newNode() *node {
	return &node{subs: make(map[string]*node)}
}

func (n *node) add(route []string, value string) bool {
	if len(route) == 0 {
		if n.value != "" {
			return false // have a conflict
		}

		n.value = value
		return true
	}

	cur := route[0]
	next, ok := n.subs[cur]
	if !ok {
		next = newNode()
		n.subs[cur] = next
	}
	return next.add(route[1:], value)
}

func (n *node) findSub(route []string) (int, string) {
	cur := route[0]
	next, ok := n.subs[cur]
	if !ok {
		return 0, ""
	}
	ret, v := next.find(route[1:])
	if v == "" {
		return 0, ""
	}
	return ret + 1, v
}

func (n *node) find(route []string) (int, string) {
	if len(route) == 0 {
		return 0, n.value
	}

	ret, v := n.findSub(route)
	if v != "" {
		return ret, v
	}

	if n.value != "" {
		return 0, n.value
	}
	return 0, ""
}
