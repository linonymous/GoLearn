package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"
)

type augmentReader struct {
	innerReader io.Reader
	augmentFunc func([]byte) []byte
}

func (r *augmentReader) Read(buf []byte) (int, error) {
	tmpBuf := make([]byte, len(buf))
	n, err := r.innerReader.Read(tmpBuf)
	copy(buf[:n], r.augmentFunc(tmpBuf[:n]))
	return n, err
}

func BangReader(r io.Reader) io.Reader {
	return &augmentReader{
		innerReader: r,
		augmentFunc: bangify,
	}
}

func UpcaseReader(r io.Reader) io.Reader {
	return &augmentReader{
		innerReader: r,
		augmentFunc: bytes.ToUpper,
	}
}

func bangify(buf []byte) []byte {
	return bytes.Replace(buf, []byte(" "), []byte("!"), -1)
}

func main() {
	originalReader := strings.NewReader("This is the stuff being read.")
	augmentReader := UpcaseReader(BangReader(originalReader))
	var res []byte
	buf := make([]byte, 4)
	for {
		n, err := augmentReader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("This shouldn't work")
		}
		res = append(res, buf[:n]...)
	}
	fmt.Println(string(res))
}
