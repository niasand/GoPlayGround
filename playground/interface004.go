package main

import "fmt"

type Shaper1 interface {
	Area() float32
}

type Square1 struct {
	side float32
}

func (sq *Square) Area2() float32 {
	return sq.side * sq.side
}

func main12() {
	sq1 := new(Square1)
	sq1.side = 5
	areainterface = sq1
	fmt.Printf("The square has area :%f\n", areainterface.Area())
}
