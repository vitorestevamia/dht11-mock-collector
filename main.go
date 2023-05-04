package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, `{"temperature": %s, "humidity": %s}`, strconv.Itoa(rand.Intn(40-20)+20), strconv.Itoa(rand.Intn(60-10)+10))
	})

	http.ListenAndServe(":5000", nil)
}
