package main

import "fmt"

type Person struct {
	Name    string
	Age     uint8
	Address Addrs
}

type Addrs struct {
	city     string
	district string
}

func main() {
	personChan := make(chan Person, 1)
	p1 := Person{"Jerry", 18, Addrs{"SZ", "Nanshan"}}
	fmt.Printf("%v\n", p1)

	personChan <- p1
	p1.Address.district = "shijiedazhan"
	fmt.Printf("%v", p1)
	p1_copy := <-personChan
	fmt.Printf("p1_copy: %v", p1_copy)
	// 通道中的元素值丝毫没有受到外界的影响。
	// 这说明了，在发送过程中进行的元素值属于完全复制。这也保证了我们使用通道传递的值的不变性。
}
