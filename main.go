package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, `{"temperature": %v, "humidity": %v}`, rand.Intn(40-20)+20, rand.Intn(60-10)+10)
	})

	http.ListenAndServe(":5000", nil)
}
