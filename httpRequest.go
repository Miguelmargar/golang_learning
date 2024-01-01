package main

import (
	"fmt"
	"net/http"
	"os"
)

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
