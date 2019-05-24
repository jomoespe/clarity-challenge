// Copyright 2018 Jomoespe. All rights reserved.
// Use of this source code is governed by a WTFPL-style
// license that can be found in the LICENSE file.package logparser_test

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
	}
	for _, test := range tests {
		if _, err := logparser.CreateReader(test.filenames...); (err == nil) == test.expectedError {
			t.Errorf("Unexpected error creating reader. Expected error=%t, Got: %t", test.expectedError, (err == nil))

		}
	}

}
