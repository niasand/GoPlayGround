package main

import "fmt"

type Simpler interface {
	Get() int
	Set(value int) int
}

type Person struct {
	age  int
	name string
}

func (p Person) Get() int {
	return p.age
}

func (p *Person) Set(value int) int {
	p.age = value
	return p.age
}

func ReadWriteAge(person Simpler, value int) {
	fmt.Printf("Get age value %v\n", person.Get())
	fmt.Printf("Set age value %v\n", person.Set(value))
	fmt.Printf("Get age value %v\n", person.Get())
}

func main() {
	var jim = &Person{12, "jim"}

	ReadWriteAge(jim, 1)

	//ReadWriteAge(jim, 13)

}
