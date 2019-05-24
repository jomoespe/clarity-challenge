// Copyright 2018 Jomoespe. All rights reserved.
// Use of this source code is governed by a WTFPL-style
// license that can be found in the LICENSE file.

package logparser_test

import (
	"testing"

	"github.com/jomoespe/clarity-challenge/pkg/logparser"
)

func TestParseLogLine(t *testing.T) {
	tests := []struct {
		logline   string
		timestamp int64
		source    string
		target    string
		err       error
	}{
		{"1234 source target", int64(1234), "source", "target", nil},
		{"1234 source target\n ", int64(1234), "source", "target", nil},
		{"1234 source", int64(1234), "source", "target", logparser.ErrNotEnoughFields},
		{"1234X source", int64(1234), "source", "target", logparser.ErrParsingDate},
	}
	for _, test := range tests {
		line, err := logparser.ParseLogLine(test.logline)
		if err != nil {
			if err != test.err {
				t.Errorf("Unexpected error parsing line. %v", err)
			}
			break
		}
		if line.Timestamp != test.timestamp {
			t.Errorf("Wrong timestamp. Expected %d, Got: %d", test.timestamp, line.Timestamp)
		}
		if line.Source != test.source {
			t.Errorf("Wrong source host. Expected '%s', Got: '%s'", test.source, line.Source)
		}
		if line.Target != test.target {
			t.Errorf("Wrong target host. Expected '%s', Got: '%s'", test.target, line.Target)
		}
	}
}
