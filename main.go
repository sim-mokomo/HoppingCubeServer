package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello,World\n")
}

func testEntryCreatedStage(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "GET" {
		fmt.Fprint(w, "test get request response")
		return
	}

	if method == "POST" {
		fmt.Fprint(w, "test post request response")
		return
	}
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/create_stage", testEntryCreatedStage)
	http.HandleFunc("/", hello)
	http.ListenAndServe(":"+port, nil)
}
