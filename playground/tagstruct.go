package main

import (
	"fmt"
	"reflect"
)

type TagType struct {
	name string "My name is jim"
	age  int    "age is 24"
	sex  string "female or male"
}

func main() {
	tt := TagType{"Lucy", 18, "Male"}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}
