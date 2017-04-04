package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go Runs On ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "Linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}
