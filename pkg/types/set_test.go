// Copyright 2018 Jomoespe. All rights reserved.
// Use of this source code is governed by a WTFPL-style
// license that can be found in the LICENSE file.

package types_test

import (
	"testing"

	"github.com/jomoespe/clarity-challenge/pkg/types"
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
		set := types.Set{}
		set.Add(test.values...)
		if len(set) != test.expected {
			t.Errorf("wrong set size. Expected: %d, Got: %d", test.expected, len(set))
		}
	}
}

func TestClean(t *testing.T) {
	tests := []struct {
		values   []string
		expected int
	}{
		{values: []string{}, expected: 0},
		{values: []string{"one", "two", "three"}, expected: 0},
		{values: []string{"one", "two", "two", "three", "two", "three"}, expected: 0},
	}

	for _, test := range tests {
		set := types.Set{}
		set.Add(test.values...)
		set.Clean()
		if len(set) != test.expected {
			t.Errorf("set not empty after Clean(). Expected: %d, Got: %d", test.expected, len(set))
		}
	}
}
