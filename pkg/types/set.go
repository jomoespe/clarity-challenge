// Copyright 2018 Jomoespe. All rights reserved.
// Use of this source code is governed by a WTFPL-style
// license that can be found in the LICENSE file.

package types

type Set map[string]struct{}

// Add values to a Set
func (set Set) Add(values ...string) {
	for _, value := range values {
		if _, ok := set[value]; !ok {
			set[value] = struct{}{}
		}
	}
}

// Clean removes all Set elements
func (set Set) Clean() {
	for k := range set {
		delete(set, k)
	}
}
