package GoExamplesReadOnly

import "fmt"

type Integers int

type LessAdder interface {
	Less(b Integers) bool
	Add(b Integers)
}

func (a Integers) Less(b Integers) bool {
	return a < b
}

func (a *Integers) Add(b Integers) {
	*a +=b
}

func mainfalse(){
	var a Integers = 1

	var lessadder LessAdder = &a

	fmt.Println(lessadder.Less(2))
}