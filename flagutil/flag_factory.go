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

package flagutil

import (
	"flag"
)

// FlagFactory is a factory that can be used to create flag sets.
type FlagFactory struct {
	name string
}

// NewFactory creates a new flag factory with the given name.
func NewFactory(name string) *FlagFactory {
	return &FlagFactory{name: name}
}

// ParseArgs parses the args using the given FlagSet, and returns
// the parsed args that does not include flags.
func ParseArgs(s *flag.FlagSet, args []string) []string {
	s.Parse(args)
	return s.Args()
}

// New creates a new flag set.
func (f *FlagFactory) New() *FlagSet {
	set := flag.NewFlagSet(f.name, flag.ExitOnError)
	return &FlagSet{
		FlagSet: set,
	}
}

// PlainArgs parse the args with no flags.
func (f *FlagFactory) PlainArgs(args []string) []string {
	flags := f.New()
	return flags.ParseArgs(args)
}

// FlagSet extends the *flag.FlagSet with some more common helper functions.
type FlagSet struct {
	*flag.FlagSet
}

// ParseArgs parses the args and returns the parsed args that does not include
// flags.
func (s *FlagSet) ParseArgs(args []string) []string {
	return ParseArgs(s.FlagSet, args)
}
