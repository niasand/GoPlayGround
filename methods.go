package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func (v *Vertex) Abshaha() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat) ABS() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

type MyInt int

func (k MyInt) aaa() int {
	if k == 0 {
		return int(k)
	}
	return int(-k)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abshaha())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.ABS())

	k := MyInt(1)
	fmt.Println(k.aaa())

}
