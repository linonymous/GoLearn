package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Create reader & writer
	r := strings.NewReader("This is the stuff being read\n")
	w := os.Stdout

	// create pipereader & pipewriter
	pr, pw := io.Pipe()
	defer pr.Close()
	defer pw.Close()

	// fork a goroutine, pipereader and pipwriter works synchronously,
	// if executed serially, there is no reader when the stuff is being
	// written. Thus fork a goroutine and then parallely write stuff when
	// it is being read :)
	go func() {
		for {
			// buf := make([]byte, 5)
			// n, err := r.Read(buf)
			n, err := io.Copy(pw, r)
			if n == 0 || err != nil {
				pw.Close() // closing is important, otherwise PipeWriter keeps on polling
				return
			}
			// pw.Write(buf)
		}
	}()

	// Read from pipeReader
	_, err := io.Copy(w, pr)
	if err != nil {
		fmt.Println(err.Error())
	}
}
