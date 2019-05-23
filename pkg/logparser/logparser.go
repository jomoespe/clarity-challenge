package logparser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Logline struct {
	Date   int64
	Source string
	Target string
}

func ParseLogLine(line string) (*Logline, error) {
	s := strings.Split(line, " ")
	if len(s) < 3 {
		return &Logline{}, errors.New("log line does not have at least three fields")
	}
	d, err := strconv.ParseInt(s[0], 10, 64)
	if err != nil {
		return &Logline{}, fmt.Errorf("error parsing date. line: %v", line)
	}
	return &Logline{Date: d, Source: s[1], Target: s[2]}, nil
}