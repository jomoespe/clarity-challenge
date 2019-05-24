// Copyright 2018 Jomoespe. All rights reserved.
// Use of this source code is governed by a WTFPL-style
// license that can be found in the LICENSE file.

package logparser

import (
	"bufio"
	"os"
)

// CreateReader for first file array element, or from standard input
// if array is empty.
func CreateReader(files ...string) (*bufio.Reader, error) {
	if len(files) < 1 {
		return bufio.NewReader(os.Stdin), nil
	}
	file, err := os.Open(files[0])
	if err != nil {
		return nil, err
	}
	return bufio.NewReader(file), nil
}
