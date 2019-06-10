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
	"time"

	"github.com/jomoespe/clarity-challenge/pkg/encoding/clarity"
	"github.com/jomoespe/clarity-challenge/pkg/logparser"
	"github.com/jomoespe/clarity-challenge/pkg/types"
)

type config struct {
	filenames          []string
	startDate, endDate time.Time
	hostname           string
	verbose            bool
}

func main() {
	c := createConfig()
	r, err := logparser.CreateReader(c.filenames...)
	if err != nil {
		fmt.Fprintln(os.Stderr, "File not found")
		os.Exit(logparser.FileNotFoundExitCode)
	}

	hosts := processLog(c, r)
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
		startDate: time.Unix(*start, 0),
		endDate:   time.Unix(*end, 0),
		hostname:  *hostname,
		verbose:   *verbose,
	}
}

func processLog(c *config, r io.Reader) *types.Set {
	found := &types.Set{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := clarity.Logline{}
		err := line.UnmarshalText([]byte(scanner.Text()))
		if err != nil {
			if c.verbose {
				fmt.Fprintf(os.Stderr, "%v\n", err)
			}
			continue
		}
		if line.Timestamp.Equal(c.startDate) || line.Timestamp.After(c.startDate) && line.Timestamp.Before(c.endDate) {
			switch {
			case c.hostname == "*":
				found.Add(line.Target, line.Source)
			case line.Source == c.hostname:
				found.Add(line.Target)
			case line.Target == c.hostname:
				found.Add(line.Source)
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
