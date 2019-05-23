package main

import (
	"fmt"
	"net/http"
)

func main() {
	url1 := "http://www.google.com"
	fmt.Printf("GET %q\n", url1)
	resp1, err := http.Get(url1)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	defer resp1.Body.Close()
	line1 := resp1.Proto + " " + resp1.Status
	fmt.Printf("First line: \n%s\n", line1)
}
