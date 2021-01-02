package GoExamplesReadOnly

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
