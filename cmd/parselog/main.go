// Copyright 2019 Jomoespe. All rights reserved.
// Use of this source code is governed by a WTFPL-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/jomoespe/clarity-challenge/pkg/encoding/clarity"
	"github.com/jomoespe/clarity-challenge/pkg/logparser"
	"github.com/jomoespe/clarity-challenge/pkg/types"
)

const (
	// OutputTriggerDuration is the duration the script will output
	OutputTriggerDuration = 1 * time.Second
)

type config struct {
	filenames []string
	host      string
	lapse     time.Duration
}

var (
	conf      = createConfig()
	senders   = types.Set{}
	receivers = types.Set{}
	conns     = types.HostConnections{}
)

func main() {
	r, err := logparser.CreateReader(conf.filenames...)
	if err != nil {
		fmt.Fprintln(os.Stderr, "File not found")
		os.Exit(logparser.FileNotFoundExitCode)
	}

	loglines := readLog(r)
	quit := processLoglines(loglines)
	<-quit
	printReport(senders, receivers)
}

func createConfig() *config {
	host := flag.String("host", "", "the host to find")
	lapse := flag.Int("lapse", 3600, "Number of seconds to generate report")
	flag.Parse()

	files := flag.Args()
	if len(files) < 1 {
		files = []string{}
	}

	return &config{
		filenames: files,
		host:      *host,
		lapse:     time.Duration(*lapse) * time.Second,
	}
}

func readLog(r *bufio.Reader) chan *clarity.Logline {
	log := make(chan *clarity.Logline)
	go func() {
		for {
			line, err := r.ReadBytes('\n')
			if err != nil {
				close(log)
				break
			}
			l := &clarity.Logline{}
			if err := l.UnmarshalText(line); err != nil {
				log <- l
			}
		}
	}()
	return log
}

func processLoglines(log chan *clarity.Logline) <-chan struct{} {
	quit := make(chan struct{})
	ticker := time.NewTicker(conf.lapse)
	go func() {
		for {
			select {
			case line, more := <-log:
				if !more {
					ticker.Stop()
					quit <- struct{}{}
					break
				}

				conns.Add(line.Source)
				if line.Source == conf.host {
					receivers.Add(line.Target)
				}
				if line.Target == conf.host {
					senders.Add(line.Source)
				}
			case <-ticker.C:
				printReport(senders, receivers)
				// clean data
				senders.Clean()
				receivers.Clean()
				conns.Clean()
			}
		}
	}()
	return quit
}

func printReport(senders, receivers types.Set) {
	now := time.Now().Unix()

	fmt.Printf("\n== Report (%v) ==================================================================\n", time.Unix(now, 0))
	fmt.Printf(" > Hosts Ccnnected to %s ________\n", conf.host)
	for host := range senders {
		fmt.Printf("\t%s\n", host)
	}
	fmt.Printf(" > Receive connections from %s _______\n", conf.host)
	for host := range receivers {
		fmt.Printf("\t%s\n", host)
	}
	host, max := conns.Max()
	fmt.Printf("\n > Host that generate max connections connections is %s (%d connectios)\n", host, max)
	fmt.Println("=====================================================================================")
}
