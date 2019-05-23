package logparser

import (
	"bufio"
	"os"
)

// CreateReader for first filename array element, or from standard input
// if array is empty.
func CreateReader(filenames ...string) (reader *bufio.Reader, err error) {
	if len(filenames) < 1 {
		reader = bufio.NewReader(os.Stdin)
	} else {
		var file *os.File
		if file, err = os.Open(filenames[0]); err == nil {
			reader = bufio.NewReader(file)
		}
	}
	return
}
