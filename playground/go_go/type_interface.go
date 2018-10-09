package main

import "fmt"

type Sayer interface {
	Say(message string)
	SayHi()
}

type Animal struct {
	Name string
}

func (a *Animal) Say(message string) { //comment
	fmt.Printf("Animal[%+v] say:%+v\n", a.Name, message)
}

func (b *Animal) SayHi() {
	b.Say("Hi,implement SayHi")
}

type Dog struct {
	Animal
}

func (d *Dog) Say(message string) {
	fmt.Printf("Dog[%v] say: %v\n", d.Name, message)
}
func main() {
	var sayer Sayer
	sayer = &Dog{Animal{Name: "Yoda"}}
	sayer.Say("Hello WOrld!")
	sayer.SayHi()

}
