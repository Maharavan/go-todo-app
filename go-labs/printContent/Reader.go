package main

import (
	"fmt"
	"io"
	"os"
)

type content struct {
	filename string
	data     []byte
	offset   int
}

func Reader() {
	ch := &content{}
	ch.filename = os.Args[1]
	fmt.Println("Filename:%s", ch.filename)
	io.Copy(os.Stdout, ch)
}

func (ch *content) Read(b []byte) (int, error) {
	if ch.data == nil {
		out, err := os.ReadFile(ch.filename)
		if err != nil {
			panic(err)
		}
		ch.data = out
	}
	if ch.offset >= len(ch.data) {
		return 0, io.EOF
	}
	n := copy(b, ch.data[ch.offset:])
	ch.offset += n
	return n, nil
}
