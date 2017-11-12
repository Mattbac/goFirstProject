package main

import (
	"net/http"
	"strconv"
)

var nbClient = 0

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	nbClient++
	var st = "<p>hello !</p><p>You are the " + strconv.Itoa(nbClient) + "</p>"
	w.Write([]byte(st))
}
