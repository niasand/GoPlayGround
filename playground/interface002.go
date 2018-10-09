package main

import (
	"fmt"
)

type bycicle struct {
	name  string
	wheel int
	color string
	price float32
}

func (s bycicle) getValue() float32 {
	return s.price
}

type car struct {
	name  string
	wheel int
	color string
	price float32
}

func (c car) getValue() float32 {
	return c.price * float32(c.wheel)
}

type valueable interface {
	getValue() float32
}

func showValue(asset valueable) {
	fmt.Printf("Value of the asset is %f\n", asset.getValue())
}

func main() {
	o := car{"BMW", 4, "Red", 340000.0}
	p := bycicle{"BMW", 4, "Red", 340000.0}

	fmt.Printf("%+v\n", o.getValue())
	fmt.Printf("%v", p.getValue())

	var q valueable = car{"BMW", 4, "Red", 340000.0}
	showValue(q)
}
