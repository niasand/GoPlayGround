package main

import (
	"fmt"
)

func main() {
	fmt.Println("\xe4\x00\x00\xe7\x95\x8cabc")

	for i, c := range "\xe4\x00\x00\xe7\x95\x8cabc" {
		fmt.Println(i, c)
	}

	for j, d := range "世界杯abc" {
		fmt.Println(j, d)
	}

	fmt.Println(str2bytes("hahaha"))

	var a []int
	a = append(a, []int{1, 2, 3, 4}...)
	fmt.Println(a)

	var c = []byte{'A', 1, 2, 3}

	b := Filter(c, ReturnBool)
	fmt.Println("b is: ", b)
}

func str2bytes(s string) []byte {
	p := make([]byte, len(s))

	for i := 0; i < len(s); i++ {
		c := s[i]
		p[i] = c
	}

	return p
}

func ReturnBool(x interface{}) bool {
	//switch x.(type) {
	//case byte:
	//	return true
	//case int:
	//	return false
	//default:
	//	return false
	//}
	if _, ok := x.(byte); ok {
		return true
	} else {
		return false
	}

}

func Filter(s []byte, fn func(x interface{}) bool) []byte {
	b := s[:0]
	for _, x := range s {
		if !fn(x) {
			b = append(b, x)
		}
	}

	return b
}
