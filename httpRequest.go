package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

// Implementing the Writer interface
func (logWriter) Write(bs []byte) (int, error) {

	fmt.Println(string(bs))

	fmt.Println("Number of bytes", len(bs))

	return len(bs), nil
}

func runCustomLogger() {

	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func runHttpRequest() {

	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Read does not have the ability to resize slices so have to give a
	// big enough empty slice to fit everything
	bs := make([]byte, 99999)

	resp.Body.Read(bs)
}

func runHttpRequestWithWrite() {

	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, resp.Body)
}
