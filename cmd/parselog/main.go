package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jomoespe/clarity-challenge/pkg/set"
)

type logline struct {
	date   int64
	source string
	target string
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

type config struct {
	filename           string
	startDate, endDate int64
	hostname           string
	verbose            bool
}

func main() {
	// parse cmd line
	config := &config{
		"test/input-file-10000.txt",
		int64(1565647204351),
		int64(1565687511867),
		"Aadvik",
		false,
	}
	processLog(config)
}

func processLog(config *config) {
	found := set.Set{}

	file, err := os.Open(config.filename)
	if err != nil {
		panic(fmt.Sprintf("error opening %s: %v", config.filename, err))
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if line, err := parse(scanner.Text()); err != nil {
			if config.verbose {
				fmt.Fprintf(os.Stderr, "%v\n", err)
			}
		} else {
			if line.date >= config.startDate && line.date <= config.endDate {
				switch {
				case line.source == config.hostname:
					found.Add(line.target)
				case line.target == config.hostname:
					found.Add(line.source)
				}
			}
		}
	}
	for host := range found {
		fmt.Println(host)
	}
}
