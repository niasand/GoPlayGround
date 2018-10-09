package main

import "fmt"

type B struct {
	name string
}

func (b B) String() string {
	return fmt.Sprintf("My name is %s", b.name)
}

func f(i interface{}) {
	if _, ok := i.(B); ok {
		//if a is B  do something
		fmt.Println("i is B let us do something")
	} else {
		fmt.Println("i is not B")
	}
}

func main() {
	var i interface{} = "yang"
	s := i.(string)
	fmt.Println(s)
	f(B{name: "foo"})
	f("hahah")

	r, ok := i.(float64)
	fmt.Println(r, ok)
}
