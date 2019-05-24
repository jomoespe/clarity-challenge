// Copyright 2018 Jomoespe. All rights reserved.
// Use of this source code is governed by a WTFPL-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"os"

	"github.com/jomoespe/clarity-challenge/pkg/logparser"
	"github.com/jomoespe/clarity-challenge/pkg/types"
)

type config struct {
	filenames          []string
	startDate, endDate int64
	hostname           string
	verbose            bool
}

func main() {
	config := createConfig()
	reader, err := logparser.CreateReader(config.filenames...)
	if err != nil {
		fmt.Fprintln(os.Stderr, "File not found")
		os.Exit(logparser.FileNotFoundExitCode)
	}

	hosts := processLog(config, reader)
	print(hosts)
}

func createConfig() *config {
	start := flag.Int64("start", 0, "start init time")
	end := flag.Int64("end", math.MaxInt64, "end init time")
	hostname := flag.String("host", "*", " the hostname")
	verbose := flag.Bool("v", false, "print errors in standard err")
	flag.Parse()

	files := flag.Args()
	if len(files) < 1 {
		files = []string{}
	}

	return &config{
		filenames: files,
		startDate: *start,
		endDate:   *end,
		hostname:  *hostname,
		verbose:   *verbose,
	}
}

func processLog(config *config, reader io.Reader) *types.Set {
	found := &types.Set{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		if line, err := logparser.ParseLogLine(scanner.Text()); err != nil {
			if config.verbose {
				fmt.Fprintf(os.Stderr, "%v\n", err)
			}
		} else {
			if line.Timestamp >= config.startDate && line.Timestamp <= config.endDate {
				switch {
				case config.hostname == "*":
					found.Add(line.Target)
					found.Add(line.Source)
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

func print(hosts *types.Set) {
	for host := range *hosts {
		fmt.Println(host)
	}
}
