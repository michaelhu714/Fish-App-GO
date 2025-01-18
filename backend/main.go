package main

import (
	"fmt"
	"github.com/michaelhu714/Fish-App-GO/api"
	"log"
	"net/http"
)

const portNum string = ":8000"

func main() {
	http.HandleFunc("/", api.Home)
	http.HandleFunc("/new", api.New)
	fmt.Printf("Server started on port %s\n", portNum)
	err := http.ListenAndServe(portNum, nil)
	if err != nil {
		log.Fatal(err)
	}
}
