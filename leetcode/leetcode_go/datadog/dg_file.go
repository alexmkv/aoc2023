package main

import (
	"fmt"
	"io"
	"os"
)

type BufWriter struct {
	wr       io.Writer
	buf      []byte
	bufBytes int
	bufSize  int
}

func (w *BufWriter) Write(p []byte) (n int, err error) {
	if len(p)+w.bufBytes <= w.bufSize {
		copy(w.buf[w.bufBytes:], p)
		w.bufBytes += len(p)
		return len(p), nil
	}
	err = w.Flush()
	if err != nil {
		return 0, err
	}
	if len(p) >= w.bufSize {
		return w.wr.Write(p)
	} else {
		return w.Write(p)
	}
}

func (w *BufWriter) Flush() error {
	if w.bufBytes == 0 {
		return nil
	}
	n, err := w.wr.Write(w.buf[:w.bufBytes])
	if n != w.bufBytes {
		copy(w.buf[0:], w.buf[n:w.bufBytes-n])
		return err
	}
	w.bufBytes = 0
	return nil
}

func NewBufWriter(wr io.Writer, bufSize int) *BufWriter {
	return &BufWriter{wr: wr, bufSize: bufSize, buf: make([]byte, bufSize)}
}

type BWriter interface {
	io.Writer
	Flush() error
}

func main() {

	f, _ := os.OpenFile("o.txt", os.O_CREATE|os.O_WRONLY, 0644)
	fi, _ := os.Open("t3016.go")
	bw := NewBufWriter(f, 10)
	buf := make([]byte, 20)

	b, _ := fi.Read(buf[0:12])
	bw.Write(buf[0:12])
	b, _ = fi.Read(buf[0:8])
	bw.Write(buf[0:8])
	b, _ = fi.Read(buf[0:3])
	bw.Write(buf[0:3])
	b, _ = fi.Read(buf[0:10])
	bw.Write(buf[0:10])
	fmt.Println(b)
}
