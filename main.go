package main

import (
	"fmt"
	"net/http"
)

//If the server is up, start client, if not, start server
//Ensures same main can be used for client and server
func main() {
	resp, err := http.Get("http://localhost:8080/courses")
	if err == nil {
		mainClient()
	} else {
		mainServer()
	}

	fmt.Println(resp.StatusCode)
}
