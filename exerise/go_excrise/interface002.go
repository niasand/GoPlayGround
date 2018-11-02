package main

import "fmt"

type Integer int

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

func (a Integer) Less(b Integer) bool {
	return a < b
}

func (a *Integer) Add(b Integer) {
	*a +=b
}

func main(){
	var a Integer = 1

	var lessadder LessAdder = &a

	fmt.Println(lessadder.Less(2))
}