package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

//
// ▶ curl -i localhost:3000
// HTTP/1.1 200 OK
// Date: Thu, 04 Jan 2018 06:37:06 GMT
// Content-Length: 2
// Content-Type: text/plain; charset=utf-8
