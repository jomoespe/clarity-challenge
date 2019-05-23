package main

import (
	"bufio"
	"fmt"
	"flag"
	"os"
	"math"
	"io"

	"github.com/jomoespe/clarity-challenge/pkg/set"
	"github.com/jomoespe/clarity-challenge/pkg/logparser"
)

type config struct {
	filenames           []string
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
	for host := range *hosts {
		fmt.Println(host)
	}
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
		endDate: *end,
		hostname: *hostname,
		verbose: *verbose,
	}
}

func processLog(config *config, reader io.Reader) *set.Set {
	found := &set.Set{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		if line, err := logparser.ParseLogLine(scanner.Text()); err != nil {
			if config.verbose {
				fmt.Fprintf(os.Stderr, "%v\n", err)
			}
		} else {
			if line.Date >= config.startDate && line.Date <= config.endDate {
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
