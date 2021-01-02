package GoExamplesReadOnly

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}

	c <- sum //将 sum 送入c
}

func mainfalse() {
	a := []int{7, 2, 8, -9, 4, 0}
	fmt.Println(a)
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}
