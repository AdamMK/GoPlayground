package main

import (
	"io"
	"log"
	"os"
)

func main() {

	f := os.Args[1]

	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
