package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, `{"temperature": 10, "humidity": 68}`)
	})

	http.ListenAndServe(":5000", nil)
}
