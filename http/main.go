package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.ca")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// hold the data from the read function
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)

	// fmt.Println(string(bs))

	// This is the same thing?!
	// io.Copy(os.Stdout, resp.Body)

	// Custom implementation of writer interface
	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes", len(bs))
	return len(bs), nil
}
