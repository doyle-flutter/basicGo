package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("JamesssssDev"))
	})
	http.HandleFunc("/html", func(res http.ResponseWriter, req *http.Request) {
		dev := "JamesssDev + GO Lang"
		html := `
		<html>
		<head>
			<title>GO HTTP Server</title>
			<script type="text/javascript" src="/assets/james.js"></script>
			<link href="/assets/james.css" rel="stylesheet" />
		</head>
		<body>
			<h2>` + dev + `</h2>
		</body>
		</html>
		`
		res.Header().Set("Content-Type", "text/html") // HTML 헤더 설정
		res.Write([]byte(html))
	})
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":8000", nil)
}
