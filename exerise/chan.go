package main

import "fmt"

func main() {
	defer start()()
	fmt.Println("should be in the middle ")
}

func start() func() {
	fmt.Println("start")
	return func() {
		fmt.Println("stop")
	}
}
