package main

import (
	"net/http"
	"path"
)

func main() {
	http.HandleFunc("/", foooo)
	http.ListenAndServe(":3000", nil)
}

func foooo(w http.ResponseWriter, r *http.Request) {
	// Assuming you want to serve a photo at 'images/foo.png'
	fp := path.Join("images", "foo.png")
	http.ServeFile(w, r, fp)

}

// ▶ curl -I localhost:3000
// HTTP/1.1 200 OK
// Accept-Ranges: bytes
// Content-Length: 161551
// Content-Type: image/png
// Last-Modified: Thu, 04 Jan 2018 07:07:59 GMT
// Date: Thu, 04 Jan 2018 07:09:33 GMT
