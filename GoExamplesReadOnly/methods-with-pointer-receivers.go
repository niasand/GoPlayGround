package GoExamplesReadOnly

import (
	"fmt"
	"math"
)

type Vertexss struct {
	X, Y float64
}

func (v *Vertexss) Abs() float64 {

	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertexss) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Receivers() {
	v := &Vertexss{3, 4}
	fmt.Println("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Println("Before scaling: %+v, Abs: %v\n", v, v.Abs())
}
