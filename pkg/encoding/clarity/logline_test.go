// Copyright 2019 Jomoespe. All rights reserved.
// Use of this source code is governed by a WTFPL-style
// license that can be found in the LICENSE file.

package clarity_test

import (
	"testing"
	"time"

	"github.com/jomoespe/clarity-challenge/pkg/encoding/clarity"
)

func TestUnmarshalText(t *testing.T) {

	tests := []struct {
		line      []byte
		timestamp time.Time
		source    string
		target    string
		err       error
	}{
		{[]byte("1234 source target"), time.Unix(int64(1234), 0), "source", "target", nil},
		{[]byte("1234 source target\n "), time.Unix(int64(1234), 0), "source", "target", nil},
		{[]byte("1234 source"), time.Unix(int64(1234), 0), "source", "target", clarity.ErrNotEnoughFields},
		{[]byte("1234X source"), time.Unix(int64(1234), 0), "source", "clarity", clarity.ErrParsingDate},
	}
	for _, test := range tests {
		line := clarity.Logline{}
		if err := line.UnmarshalText(test.line); err != nil {
			if err != test.err {
				t.Errorf("Unexpected error parsing line. %v", err)
			}
			break
		}
		if line.Timestamp != test.timestamp {
			t.Errorf("Wrong timestamp. Expected %T, Got: %T", test.timestamp, line.Timestamp)
		}
		if line.Source != test.source {
			t.Errorf("Wrong source host. Expected '%s', Got: '%s'", test.source, line.Source)
		}
		if line.Target != test.target {
			t.Errorf("Wrong target host. Expected '%s', Got: '%s'", test.target, line.Target)
		}
	}
}
