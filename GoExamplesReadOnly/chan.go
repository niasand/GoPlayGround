package GoExamplesReadOnly

import "fmt"

func mainfalse() {
	defer start()()
	fmt.Println("should be in the middle ")
}

func start() func() {
	fmt.Println("start")
	return func() {
		fmt.Println("stop")
	}
}
