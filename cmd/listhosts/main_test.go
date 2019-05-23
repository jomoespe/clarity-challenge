package main

import "testing"

var ProcessLog = processLog

func TestProcessLog(t *testing.T) {
	tests := []struct {
		config   *config
		expected int
	}{
		{config: &config{
			"../../test/input-file-10000.txt",
			int64(1565647204351),
			int64(1565687511867),
			"Aadvik",
			false,
		},
			expected: 3},
	}
	for _, test := range tests {
		found := *processLog(test.config)
		if len(found) != test.expected {
			t.Errorf("wrong number of elements processed size. Expected: %d, Got: %d", test.expected, len(found))
		}
	}
}
