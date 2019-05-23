package set_test

import (
	"testing"

	"github.com/jomoespe/clarity-challenge/pkg/set"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		values   []string
		expected int
	}{
		{values: []string{"one", "two", "three"}, expected: 3},
		{values: []string{"one", "two", "two", "three", "two", "three"}, expected: 3},
	}

	for _, test := range tests {
		set := set.Set{}
		set.Add(test.values...)
		if len(set) != test.expected {
			t.Errorf("wrong set size. Expected: %d, Got: %d", test.expected, len(set))
		}
	}
}
