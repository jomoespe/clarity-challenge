package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"time"
)

const (
	// FileNotFoundExitCode is the exit code when file is not found
	FileNotFoundExitCode  = 1
	
	// IOErrorExitCode is the exit when an I/O error is triggered procesing input file
	IOErrorExitCode       = 2
	
	// OutputTriggerDuration is the duration the script will output
	OutputTriggerDuration = 1 * time.Second
)

func main() {
	reader, err := reader(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "File not found")
		os.Exit(FileNotFoundExitCode)
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
		    os.Exit(IOErrorExitCode)
		}
		fmt.Print("=> " + line)
	}
	quit<- struct{}{}
	fmt.Println("Sacabó!")
}

// Creates a new reader for first filename array element, or from stdout
// if array is empty.
func reader(filenames []string) (reader *bufio.Reader, err error) {
	if len(filenames) <= 1 {
		reader = bufio.NewReader(os.Stdin)
	} else {
		var file *os.File
		if file, err = os.Open(filenames[1]); err == nil {
			reader = bufio.NewReader(file)
		}
	}
	return 
}
