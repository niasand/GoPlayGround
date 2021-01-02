package GoExamplesReadOnly

import "fmt"

type Work interface {
	doWork(n int)
}

type worker struct {
	name string
}

type number struct {
	x int
	y int
}

func (a *worker) doWork(n int) {
	fmt.Printf("i am worker %v - %v\n", n, a.name)
}
func (i *number) doWork(n int) {
	fmt.Printf("I am number %v - %v %v\n", n, i.x, i.y)
}

func magic(i Work, n int) {
	i.doWork(n)
}

func mainfalse() {
	var a1 = &worker{"Jim"}
	var b1 = &number{22, 33}
	var b2 = &number{}

	magic(a1, 22)
	magic(b1, 33)
	magic(b2, 44)
}
