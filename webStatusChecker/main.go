package main

import (
	"fmt"
	"net/http"
)

func main() {

	//each request is made one after the other
	links := []string {
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
	}

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string){
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Error - it might be down", err)
		return
	}
	fmt.Println(link, "is up!")
}
