package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var input string
	fmt.Scanln(&input)
	rsr := NewReverseStringReader(input)
	io.Copy(os.Stdout, rsr)
}

type ReverseStringReader struct {
	data []byte
	done *bool
}

func NewReverseStringReader(input string) *ReverseStringReader {
	done := false
	return &ReverseStringReader{[]byte(input), &done}
}

func (rsr ReverseStringReader) Read(p []byte) (n int, err error) {
	if *rsr.done {
		return 0, io.EOF
	}

	for i, b := range []byte(rsr.data) {
		p[len(rsr.data)-i-1] = b
	}

	*rsr.done = true
	return len(rsr.data), nil
}
