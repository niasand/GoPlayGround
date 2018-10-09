package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func main() {
	sq1 := new(Square)
	sq1.side = 5

	areaInfd := sq1
	fmt.Printf("The square has area :%f\n", areaInfd.Area())

	var areaIntef Shaper
	areaIntef = sq1
	//areaIntef = Shaper(sq1)

	fmt.Printf("The square has area :%f\n", areaIntef.Area())

}
