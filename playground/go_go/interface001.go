package main

import (
	"fmt"
	"reflect"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

func (h Human) SayHi() {
	fmt.Printf("Hi, 我的名字 %s 我的电话 %s\n", h.name, h.phone)
}

func (h Human) Sing(lyrics string) {
	fmt.Println("唱歌\n", lyrics)
}

func (e Employee) SayHi() {
	fmt.Printf("员工类:名字-%s 公司-%s 工资-%v\n", e.name, e.company, e.money)
}

type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	mike := Student{Human{"Mike", 25, "18812345678"}, "MIT", 1000}

	staff := Employee{Human{"employee", 30, "hahahah"}, "Sf-express", 3000}

	//var i Men
	mike.SayHi()
	staff.SayHi()

	var yang Men
	yang = mike
	yang.SayHi()

	t := reflect.TypeOf(yang)
	//m := reflect.ValueOf(yang)
	fmt.Println(t)

}
