package main

import (
	"fmt"
)

//一个结构体（ struct ）就是一个字段的集合。

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
