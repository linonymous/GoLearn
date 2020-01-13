package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// io.Copy uses 32kb buffer and streams data from the reader through
// the buffer and writes to the writer. Copy makes it easy
func main() {
	reader := strings.NewReader("This is the stuff I'm reading\n")
	writer := os.Stdout

	n, err := io.Copy(writer, reader)
	fmt.Printf("%d bytes written\n", n)
	if err != nil {
		log.Fatal(err)
	}

	// Reads from string and writes to a writer
	a, err := io.WriteString(os.Stdout, "test\n")
	fmt.Printf("%d bytes written\n", a)
	if err != nil {
		log.Fatal(err)
	}
}
