package main

import (
	"io"
	"os"
)

func main() {
	content, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, content)
}
