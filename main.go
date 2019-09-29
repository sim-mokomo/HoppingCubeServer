package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type StageConfig struct {
	Name        string
	PanelLayout string
}

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

		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		requestLength, err := strconv.Atoi(r.Header.Get("Content-Length"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		body := make([]byte, requestLength)
		requestLength, err = r.Body.Read(body)
		if err != nil && err != io.EOF {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var jsonBody map[string]interface{}
		err = json.Unmarshal(body[:requestLength], &jsonBody)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, jsonBody)
		w.WriteHeader(http.StatusOK)
		return
	}
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/create_stage", testEntryCreatedStage)
	http.HandleFunc("/", hello)
	http.ListenAndServe(":"+port, nil)
}
