package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["age"] = 27
	fmt.Println("The value is: ", m["age"])

	m["age"] = 28
	fmt.Println("The value is: ", m["age"])

	delete(m, "age")
	fmt.Println("The value is: ", m["age"])

	m["age"] = 29
	v, ok := m["age"]
	fmt.Println("The value:", v, "ok Present?", ok)

	n := make(map[int]int)
	n[1] = 2
	fmt.Println(n[1])
}
