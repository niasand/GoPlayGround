package main

import "fmt"

func main() {
	var value int
	var isPresent bool
	city := make(map[string]int)
	city["New York"] = 100
	city["Beijing"] = 101
	value, isPresent = city["Beijinsg"]
	fmt.Println(value, isPresent)

	st := []string{"a", "b", "c", "d", "e", "f"}
	for i, letter := range st {
		fmt.Println(i, letter)
		if i == 0 {
			st = append(st[:i], st[i+1:]...)
			fmt.Println(st)
		}

	}

}
