package main

import (
	"fmt"
	"os"
)
func main() {

	var s, sep string
	for i :=1; i< len(os.Args); i++ {

		s += sep + os.Args[i]
		sep = " "
		fmt.Print(s)
		fmt.Print(i)
		fmt.Print("\n")

	}

	fmt.Print(os.Args[0])
}
