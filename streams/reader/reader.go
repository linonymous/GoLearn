package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

/*
	Simple reader example using io.Reader
*/
func main() {
	reader := strings.NewReader("This is the stuff being read.")
	var result []byte
	buf := make([]byte, 4)
	for {
		n, err := reader.Read(buf)
		// Print the contents of buffer during read
		// fmt.Println(string(buf))
		result = append(result, buf[:n]...)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
	}
	fmt.Println(string(result))
}
