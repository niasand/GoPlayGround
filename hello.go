package main

import "fmt"

func x()  {
	var x, y  int
	fmt.Println(&x == &x,&x == &y,&x==nil)
}
var p = f()
func f() *int {
	v := 1
	return &v
	
}

func main() {
	x := 1
	p := &x
	fmt.Println(*p,&x)
	*p = 2
	fmt.Println(x,&x)
	fmt.Print("Hello\n")

	fmt.Println(f() == f())
	fmt.Println(f())

}