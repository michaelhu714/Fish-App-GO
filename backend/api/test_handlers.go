package api

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage")
}

func New(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New Page")
}
