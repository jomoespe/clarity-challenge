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
	}{
		{"1234 source target", int64(1234), "source", "target"},
		{"1234 source target\n ", int64(1234), "source", "target"},
	}
	for _, test := range tests {
		line, err := logparser.ParseLogLine(test.logline)
		if err != nil {
			t.Errorf("Unexpected error parsing line. %v", err)
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
