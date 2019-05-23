package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jomoespe/clarity-challenge/pkg/set"
	"github.com/jomoespe/clarity-challenge/pkg/logparser"
)

type config struct {
	filename           string
	startDate, endDate int64
	hostname           string
	verbose            bool
}

func main() {
	// TODO parse cmd line
	config := &config{
		"test/input-file-10000.txt",
		int64(1565647204351),
		int64(1565687511867),
		"Aadvik",
		false,
	}
	found := processLog(config)
	for host := range *found {
		fmt.Println(host)
	}
}

func processLog(config *config) *set.Set {
	file, err := os.Open(config.filename)
	if err != nil {
		panic(fmt.Sprintf("error opening %s: %v", config.filename, err))
	}

	found := &set.Set{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if line, err := logparser.ParseLogLine(scanner.Text()); err != nil {
			if config.verbose {
				fmt.Fprintf(os.Stderr, "%v\n", err)
			}
		} else {
			if line.Date >= config.startDate && line.Date <= config.endDate {
				switch {
				case line.Source == config.hostname:
					found.Add(line.Target)
				case line.Target == config.hostname:
					found.Add(line.Source)
				}
			}
		}
	}
	return found
}
