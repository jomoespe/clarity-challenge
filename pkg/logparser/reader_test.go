// Copyright 2019 Jomoespe. All rights reserved.
// Use of this source code is governed by a WTFPL-style
// license that can be found in the LICENSE file.

package logparser_test

import (
	"testing"

	"github.com/jomoespe/clarity-challenge/pkg/logparser"
)

func TestCreateReader(t *testing.T) {
	tests := []struct {
		filenames     []string
		expectedError bool
	}{
		{[]string{"../../test/input-file-10000.txt"}, false},
		{[]string{"../../test/input-file-10000.txt", "file-does-not-exist"}, false},
		{[]string{"file-does-not-exist"}, true},
		{[]string{""}, true},
		{[]string{}, false},
	}

	for _, test := range tests {
		_, err := logparser.CreateReader(test.filenames...)
		if (err == nil) == test.expectedError {
			t.Errorf("Unexpected error creating reader for %s. Expected error=%t, Got: %t", test.filenames, test.expectedError, (err == nil))
			t.Error(err)
		}
	}
}
