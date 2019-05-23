package main

import (
	"fmt"
	"os"
	"io"
	"time"

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
	
	for {
		line, err := reader.ReadString('\n')
		if err != nil{
		    if err == io.EOF {
				break
		    }
		    fmt.Fprintln(os.Stderr, err)
		    os.Exit(logparser.IOErrorExitCode)
		}
		fmt.Print("=> " + line)
	}
	quit<- struct{}{}
	fmt.Println("Sacabó!")
}
