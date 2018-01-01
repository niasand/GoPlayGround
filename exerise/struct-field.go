package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{3, 5}
	v.X = 7
	fmt.Println(Vertex{3, 5}.X)
	fmt.Println(v)
}
