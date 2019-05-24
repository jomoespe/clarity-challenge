package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/jomoespe/clarity-challenge/pkg/logparser"
	"github.com/jomoespe/clarity-challenge/pkg/set"
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

var conf = createConfig()
var senders = set.Set{}
var receivers = set.Set{}
var conns = set.HostConnections{}

func main() {
	reader, err := logparser.CreateReader(conf.filenames...)
	if err != nil {
		fmt.Fprintln(os.Stderr, "File not found")
		os.Exit(logparser.FileNotFoundExitCode)
	}

	loglines := readLog(reader)
	quit := processLoglines(loglines)
	<-quit
	printReport(senders, receivers)
}

func createConfig() *config {
	host := flag.String("host", "", "the host to find")
	lapse := flag.Int("lapse", 3600, "Number of seconds to gererate report")
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

func readLog(reader *bufio.Reader) chan *logparser.Logline {
	out := make(chan *logparser.Logline)
	go func() {
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				close(out)
				break
			}
			Logline, _ := logparser.ParseLogLine(line)
			out <- Logline
		}
	}()
	return out
}

func processLoglines(loglines chan *logparser.Logline) <-chan struct{} {
	quit := make(chan struct{})
	ticker := time.NewTicker(conf.lapse)

	go func() {
		for {
			select {
			case line, more := <-loglines:
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
			}
		}
	}()
	return quit
}

func printReport(senders, receivers set.Set) {
	fmt.Println("\n== Report =========================================================================")
	fmt.Printf(" > Connected to %s ________\n", conf.host)
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
