package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	t := time.Now()
	fmt.Fprintln(w, t.String())
}

func unix(w http.ResponseWriter, req *http.Request) {
	t := time.Now().Format(time.UnixDate)
	fmt.Fprintln(w, t)
}

func kitchen(w http.ResponseWriter, req *http.Request) {
	t := time.Now().Format(time.Kitchen)
	fmt.Fprintln(w, t)
}

func main() {

	http.HandleFunc("/", hello)
	http.HandleFunc("/kitchen", kitchen)
	http.HandleFunc("/unix", unix)

	http.ListenAndServe(":8081", nil)
}
