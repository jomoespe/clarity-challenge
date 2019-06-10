// Copyright 2019 Jomoespe. All rights reserved.
// Use of this source code is governed by a WTFPL-style
// license that can be found in the LICENSE file.

package clarity

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

type Logline struct {
	Timestamp time.Time
	Source    string
	Target    string
}

// Errors
var (
	ErrNotEnoughFields = errors.New("log line does not have at least three fields")
	ErrParsingDate     = errors.New("error parsing date")
)

func (l *Logline) UnmarshalText(line []byte) error {
	s := strings.Split(string(line), " ")
	if len(s) < 3 {
		return ErrNotEnoughFields
	}
	t, err := strconv.ParseInt(s[0], 10, 64)
	if err != nil {
		return ErrParsingDate
	}
	l.Timestamp = time.Unix(t, 0)
	l.Source = s[1]
	l.Target = s[2][:len(s[2])]
	if l.Target[len(l.Target)-1:] == "\n" {
		l.Target = l.Target[:len(l.Target)-1]
	}
	return nil
}
