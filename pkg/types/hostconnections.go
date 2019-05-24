// Copyright 2018 Jomoespe. All rights reserved.
// Use of this source code is governed by a WTFPL-style
// license that can be found in the LICENSE file.package logparser_test

package types

import (
	"sort"
)

type HostConnections map[string]int

func (c HostConnections) Add(host string) {
	i := 1
	if _, ok := c[host]; ok {
		i = c[host] + 1
	}
	c[host] = i
}

func (c HostConnections) Count(host string) int {
	if _, ok := c[host]; ok {
		return c[host]
	} else {
		return 0
	}
}

// Max returns the host with max connections and the total of connections
func (c HostConnections) Max() (string, int) {
	s := Sort(c)
	maxHost := ""
	maxConns := 0
	for k, v := range s {
		if v > maxConns {
			maxHost = k
			maxConns = v
		}
	}
	return maxHost, maxConns
}

// Sort returns a sorted copy of the map
func Sort(m HostConnections) HostConnections {
	n := map[int][]string{}
	var a []int

	for k, v := range m {
		n[v] = append(n[v], k)
	}
	for k := range n {
		a = append(a, k)
	}
	//	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	r := make(map[string]int)
	for _, k := range a {
		for _, s := range n[k] {
			r[s] = k
		}
	}
	return r
}

// Clean removes all Set elements
func (c HostConnections) Clean() {
	for k := range c {
		delete(c, k)
	}
}
