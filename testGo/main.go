package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(req.Host)
		fmt.Println(req.URL.Path)
		fmt.Println(req.URL.Query())
		res.Write([]byte("James Dev"))
	})
	http.HandleFunc("/my/", myHandle)
	http.ListenAndServe(":8000", nil)
}

func myHandle(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	fmt.Println(req.URL.Query())
	res.Write([]byte("MY - James Dev"))
}
