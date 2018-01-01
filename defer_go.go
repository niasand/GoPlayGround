package main

import "fmt"

func main()  {
	yang()

}

func yang()  {
	fmt.Printf("In yang at the top!\n")
	defer zhi()
	fmt.Printf("In yang at the bottom\n")
}

func zhi(){
	fmt.Printf("function zhi:Deferred until the end of the calling function\n")
}