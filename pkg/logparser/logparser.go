// Copyright 2018 Jomoespe. All rights reserved.
// Use of this source code is governed by a WTFPL-style
// license that can be found in the LICENSE file.

package logparser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Logline struct {
	Timestamp int64
	Source    string
	Target    string
}

func ParseLogLine(line string) (*Logline, error) {
	s := strings.Split(line, " ")
	if len(s) < 3 {
		return &Logline{}, errors.New("log line does not have at least three fields")
	}
	timestamp, err := strconv.ParseInt(s[0], 10, 64)
	if err != nil {
		return &Logline{}, fmt.Errorf("error parsing date. line: %v", line)
	}
	source := s[1]
	target := s[2][:len(s[2])]
	if target[len(target)-1:] == "\n" {
		target = target[:len(target)-1]
	}
	return &Logline{timestamp, source, target}, nil
}
