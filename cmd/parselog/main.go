package main

import (
	"fmt"
	"os"
	"time"
	"bufio"

	"github.com/jomoespe/clarity-challenge/pkg/logparser"
)

const (
	// OutputTriggerDuration is the duration the script will output
	OutputTriggerDuration = 1 * time.Second
)

func main() {
	filenames := os.Args[1:]
	reader, err := logparser.CreateReader(filenames...)
	if err != nil {
		fmt.Fprintln(os.Stderr, "File not found")
		os.Exit(logparser.FileNotFoundExitCode)
	}
	ticker := time.NewTicker(OutputTriggerDuration)
	quit := make(chan struct{})
	go func() {
		for {
		   select {
			case <- ticker.C:
				// do stuff
				fmt.Println("Tick!")
			case <- quit:
				ticker.Stop()
				return
			}
		}
	 }()
	
	processing := x(reader)
	for  {
		line, more := <- processing
		if !more {
			break
		}
		fmt.Printf("timestamp: %d source: %s target: %s", line.Timestamp, line.Source, line.Target)
	}
}

func x(reader *bufio.Reader) <-chan *logparser.Logline {
	out := make(chan *logparser.Logline)
	go func() {
		for {
			line, err := reader.ReadString('\n')
			if err != nil{
				close(out)
				break
			}
			Logline, _ := logparser.ParseLogLine(line)
			out <- Logline
		}
	}()
	return out
}