package main

import "io"

import "bytes"

import "os"

import "strings"

import "log"

type augmentReader struct {
	innerReader io.Reader
	augmentFunc func([]byte) []byte
}

func (a *augmentReader) Read(buf []byte) (int, error) {
	tmpbuf := make([]byte, len(buf))
	n, err := a.innerReader.Read(tmpbuf)
	copy(buf[:n], a.augmentFunc(tmpbuf[:n]))
	return n, err
}

func bangify(buf []byte) []byte {
	return bytes.Replace(buf, []byte(" "), []byte("!"), -1)
}

func BangReader(r io.Reader) io.Reader {
	return &augmentReader{
		innerReader: r,
		augmentFunc: bangify,
	}
}

func main() {
	reader := strings.NewReader("This is the stuff being read.")
	_, err := io.Copy(os.Stdout, BangReader(reader))
	if err != nil {
		log.Fatal(err)
	}
}
