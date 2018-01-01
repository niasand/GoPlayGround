package main

import (
	"fmt"
	"os"
)

type Reader interface {
	Read(b []byte) (n int, report_err error)
}

type Writer interface {
	Write(b []byte) (n int, report_err error)
}

type ReaderWriter interface {
	Reader
	Writer
}

func main() {
	var w Writer
	w = os.Stdout
	fmt.Fprintf(w, "Hello interface\n")
}
