package main

import "fmt"

type MyType struct {
	i int
}

type Age int

func (A Age) getAge() int{
	return int(A)  //类型转换
}

func f(a int) {

}


func (p *MyType) get() int {
	return p.i
}

func (p *MyType) sget()  {
	fmt.Println(p.i)
}

type hahahha interface {
	sget()
	set(i int)
}

func(p *MyType) set(i int){
	p.i = i
}

func getAdnSet(x hahahha){}

func main(){
	var t MyType
	t.i = 2
	fmt.Println(t.get())

	var age Age
	age =1

	f(t.i)

	fmt.Println(age.getAge())

	getAdnSet(&t)

}