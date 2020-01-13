package main

import (
	"fmt"
	"io"
	"log"
)

type myReader struct {
	content  []byte
	position int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Read function reads from content of reader and
// puts in buffer passed
func (r *myReader) Read(buf []byte) (int, error) {
	rem := len(r.content) - r.position
	n := min(rem, len(buf))
	if n == 0 {
		return 0, io.EOF
	}
	copy(buf[:n], r.content[r.position:r.position+n])
	r.position += n
	return n, nil
}

func main() {
	reader := myReader{
		content: []byte("this is the stuff being read."),
	}
	var res []byte
	buf := make([]byte, 4)
	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("Something's broken!")
		}
		res = append(res, buf[:n]...)
	}
	fmt.Println(string(res))
}
