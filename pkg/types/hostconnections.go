// Copyright 2019 Jomoespe. All rights reserved.
// Use of this source code is governed by a WTFPL-style
// license that can be found in the LICENSE file.

package types

type HostConnections map[string]int

func (c HostConnections) Add(host string) {
	i := 1
	if _, ok := c[host]; ok {
		i = c[host] + 1
	}
	c[host] = i
}

func (c HostConnections) Count(host string) int {
	if _, ok := c[host]; !ok {
		return 0
	}
	return c[host]
}

// Max returns the host with max connections and the total of connections
func (c HostConnections) Max() (string, int) {
	maxHost := ""
	maxConns := 0
	for k, v := range c {
		if v > maxConns {
			maxHost = k
			maxConns = v
		}
	}
	return maxHost, maxConns
}

// Clean removes all Set elements
func (c HostConnections) Clean() {
	for k := range c {
		delete(c, k)
	}
}
