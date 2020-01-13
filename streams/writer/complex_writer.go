package main

import "io"

import "bytes"

import "os"

type augmentWriter struct {
	innerWriter io.Writer
	augmentFunc func([]byte) []byte
}

func bangify(buf []byte) []byte {
	return bytes.Replace(buf, []byte(" "), []byte("!"), -1)
}

func (w *augmentWriter) Write(buf []byte) (int, error) {
	return w.innerWriter.Write(w.augmentFunc(buf))
}

func BangWriter(w io.Writer) io.Writer {
	return &augmentWriter{
		innerWriter: w,
		augmentFunc: bangify,
	}
}

func UpcaseWriter(w io.Writer) io.Writer {
	return &augmentWriter{
		innerWriter: w,
		augmentFunc: bytes.ToUpper,
	}
}

func main()  {
	augmentedWriter := UpcaseWriter(BangWriter(os.Stdout))
	augmentedWriter.Write([]byte("Let's see if this works\n"))
}