package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("James Dev"))
	})
	http.ListenAndServe(":8000", nil)
}
