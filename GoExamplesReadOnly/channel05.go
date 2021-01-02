package GoExamplesReadOnly

import "fmt"

// Person ...
type Person struct {
	name    string
	age     uint8
	Address Addr
}

// Addr ...
type Addr struct {
	city     string
	district string
}

func mainfalse() {
	personChan := make(chan Person, 1)

	p1 := Person{"Jerry", 17, Addr{"SZ", "NS"}}
	personChan <- p1
	fmt.Println(personChan, p1)
	// Println output: 0xc420058060 {Jerry 17 {SZ NS}}

	p1Receive := <-personChan
	fmt.Println(p1Receive)
	// output: {Jerry 17 {SZ NS}}

	fmt.Println(p1Receive.age, p1Receive.Address, p1Receive.name)
	//  output:  17 {SZ NS} Jerry
}
