// Copyright 2019 Jomoespe. All rights reserved.
// Use of this source code is governed by a WTFPL-style
// license that can be found in the LICENSE file.

package types_test

import (
	"testing"

	"github.com/jomoespe/clarity-challenge/pkg/types"
)

func TestAddHostConnection(t *testing.T) {
	conns := types.HostConnections{}
	conns.Add("host-1")
	conns.Add("host-2")
	conns.Add("host-1")
	conns.Add("host-1")
	conns.Add("host-1")
	conns.Add("host-1")
	conns.Add("host-1")

	if len(conns) != 2 {
		t.Errorf("Unexpected host connections size. Expected: 2, Got: %d", len(conns))
	}
}

func TestCountHostConnections(t *testing.T) {
	conns := types.HostConnections{}
	conns.Add("host-1")
	conns.Add("host-2")
	conns.Add("host-1")
	conns.Add("host-1")
	conns.Add("host-2")
	conns.Add("host-1")
	conns.Add("host-1")

	host1Conns := conns.Count("host-1")
	host2Conns := conns.Count("host-2")
	host3Conns := conns.Count("host-3")

	if host1Conns != 5 {
		t.Errorf("Unexpected host connections for host1. Expected: 5, Got: %d", host1Conns)
	}
	if host2Conns != 2 {
		t.Errorf("Unexpected host connections for host2. Expected: 2, Got: %d", host2Conns)
	}
	if host3Conns != 0 {
		t.Errorf("Unexpected host connections for host3. Expected: 0, Got: %d", host3Conns)
	}
}

func TestMaxHostConnections(t *testing.T) {
	conns := types.HostConnections{}
	conns.Add("host-3")
	conns.Add("host-1")
	conns.Add("host-2")
	conns.Add("host-1")
	conns.Add("host-1")
	conns.Add("host-2")
	conns.Add("host-1")
	conns.Add("host-1")

	host, count := conns.Max()
	if host != "host-1" {
		t.Errorf("Wrong max host. Expected: host1, Got: %s", host)
	}
	if count != 5 {
		t.Errorf("Wrong max host count. Expected: 5, Got: %d", count)
	}
}

func TestCleanHostConnections(t *testing.T) {
	conns := types.HostConnections{}
	conns.Add("host-3")
	conns.Add("host-1")
	conns.Add("host-2")
	conns.Add("host-1")
	conns.Add("host-1")
	conns.Add("host-2")
	conns.Add("host-1")
	conns.Add("host-1")

	conns.Clean()
	if len(conns) != 0 {
		t.Errorf("Host connections not clean\n")
	}
}
