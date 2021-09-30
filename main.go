package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8080/courses")
	if err == nil {
		mainClient()
	} else {
		mainServer()
	}

	fmt.Println(resp.StatusCode)
}
