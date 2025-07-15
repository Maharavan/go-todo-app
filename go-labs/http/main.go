package main

import (
	"fmt"
	"io"
	"net/http"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://www.google.com/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bs := logWriter{}

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
	// file, err := os.Create("output.html")
	// io.Copy(file, resp.Body) used for writing response from into file where 1st argument is writer and 2nd is reader interface
	//Reader receives bytes of slice ,reads raw content and update the same
	io.Copy(bs, resp.Body)
}

func (logWriter) Write(b []byte) (int, error) {
	fmt.Println(string(b))
	return len(b), nil
}
