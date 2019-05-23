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

func main() {
	//config := createConfig()
	reader, err := logparser.CreateReader(conf.filenames...)
	if err != nil {
		fmt.Fprintln(os.Stderr, "File not found")
		os.Exit(logparser.FileNotFoundExitCode)
	}

	loglines := readLog(reader)
	quit := processLoglines(loglines)
	<-quit
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
	senders := set.Set{}
	receivers := set.Set{}

	go func() {
		for {
			select {
			case line, more := <-loglines:
				if !more {
					ticker.Stop()
					quit <- struct{}{}
					break
				}
				senders.Add(line.Source)
				receivers.Add(line.Target)
				//				fmt.Printf("timestamp: %d source: %s target: %s", line.Timestamp, line.Source, line.Target)
			case <-ticker.C:
				// print report
				fmt.Println("__ Senders _____________________")
				for host := range senders {
					fmt.Println(host)
				}
				fmt.Println("__ Receivers __________________")
				for host := range receivers {
					fmt.Println(host)
				}
				// clean data
				senders.Clean()
				receivers.Clean()
			}
		}
	}()
	return quit
}

/*
func y(lines chan logparser.Logline) {
	ticker := time.NewTicker(OutputTriggerDuration)
	quit := make(chan struct{})
	senders := set.Set{}
	receivers := set.Set{}
	go func() {
		for {
		   select {
		   case line, more := <- lines:
			if !more {
				break
			}

			case <- ticker.C:
				// do stuff
				fmt.Println("Tick!")
			case <- quit:
				ticker.Stop()
				return
			}
		}
	 }()
}
*/
