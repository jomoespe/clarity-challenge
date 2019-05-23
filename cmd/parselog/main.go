package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type logline struct {
	date   int64
	source string
	target string
}

type set map[string]struct{}

func (set set) add(host string) {
	if _, ok := set[host]; !ok {
		set[host] = struct{}{}
	}
}

func parse(line string) (*logline, error) {
	s := strings.Split(line, " ")
	if len(s) < 3 {
		return &logline{}, errors.New("log line does not have at least three fields")
	}
	d, err := strconv.ParseInt(s[0], 10, 64)
	if err != nil {
		return &logline{}, fmt.Errorf("error parsing date. line: %v", line)
	}
	return &logline{date: d, source: s[1], target: s[2]}, nil
}

func main() {
	// parse cmd line
	filename := "test/input-file-10000.txt"
	startDate := int64(1565647204351)
	endDate := int64(1565687511867)
	hostname := "Aadvik"
	verbose := false

	found := set{}

	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("error opening %s: %v", filename, err))
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if line, err := parse(scanner.Text()); err != nil {
			if verbose {
				fmt.Fprintf(os.Stderr, "%v\n", err)
			}
		} else {
			if line.date >= startDate && line.date <= endDate {
				switch {
				case line.source == hostname:
					found.add(line.target)
				case line.target == hostname:
					found.add(line.source)
				}
			}
		}
	}
	for host := range found {
		fmt.Println(host)
	}
}
