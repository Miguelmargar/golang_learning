package main

import (
	"fmt"
	"io"
	"os"
)

func runPrintFile() {

	fn := os.Args[1]

	f, err := os.Open(fn)

	if err != nil {

		fmt.Println("Error:", err)

		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
