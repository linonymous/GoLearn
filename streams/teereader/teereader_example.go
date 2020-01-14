package main

import "strings"

import "io"

import "os"

import "fmt"

import "io/ioutil"

func main() {
	reader := strings.NewReader("This is the stuff being read!\n")

	// teereader writes to a writer and returns a reader,
	// which can be read again and printed on screen
	p := io.TeeReader(reader, os.Stdout)
	s, _ := ioutil.ReadAll(p)
	fmt.Print(string(s))
}
