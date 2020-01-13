package main

import "os"

func main() {
	writer := os.Stdout
	writer.Write([]byte("This is the stuff being written."))
}
