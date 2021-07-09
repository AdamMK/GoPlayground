package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	//create empty slice with space for 99999 spaces
	//bs := make([]byte, 99999)
	//passing bs to read function
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))

	//exactly same as above
	//first argument implements Writer Interface which writes data to when we want it
	//second argument implements Reader Interface which takes the data
	//io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}

	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error)  {
	fmt.Println(string(bs))
	fmt.Println("Lines size ", len(bs))
	return len(bs), nil
}
