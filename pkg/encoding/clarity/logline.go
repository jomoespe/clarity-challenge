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

// Logline represents a line in a log line
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

// UnmarshalText is the interface implementation of encoding/UnmarshalText .
func (l *Logline) UnmarshalText(line []byte) error {
	s := strings.Split(string(line), " ")
	if len(s) < 3 {
		return ErrNotEnoughFields
	}
	sec, err := strconv.ParseInt(s[0], 10, 64)
	if err != nil {
		return ErrParsingDate
	}
	l.Timestamp = time.Unix(sec, 0)
	l.Source = strings.TrimSpace(s[1])
	l.Target = strings.TrimSpace(s[2])
	return nil
}
