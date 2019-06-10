// Copyright 2019 Jomoespe. All rights reserved.
// Use of this source code is governed by a WTFPL-style
// license that can be found in the LICENSE file.

package main

import (
	"testing"
	"time"

	"github.com/jomoespe/clarity-challenge/pkg/logparser"
)

var ProcessLog = processLog

func TestProcessLog(t *testing.T) {
	tests := []struct {
		config   *config
		expected int
	}{
		{config: &config{
			[]string{"../../test/input-file-10000.txt"},
			time.Unix(int64(1565647204351), 0),
			time.Unix(int64(1565687511867), 0),
			"Aadvik",
			false,
		},
			expected: 3},
	}
	for _, test := range tests {
		r, _ := logparser.CreateReader(test.config.filenames...)
		found := *processLog(test.config, r)
		if len(found) != test.expected {
			t.Errorf("wrong number of elements processed size. Expected: %d, Got: %d", test.expected, len(found))
		}
	}
}
