package main

import (
	"fmt"
	"log"
	"net/http"
)

const portNum string = ":8000"

func main() {
	http.HandleFunc("/", nil)
	http.HandleFunc("/new", nil)
	fmt.Printf("Server started on port %s\n", portNum)
	err := http.ListenAndServe(portNum, nil)
	if err != nil {
		log.Fatal(err)
	}
}
