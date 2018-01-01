package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) String() string {
	return fmt.Sprintf("%v: %v years.", p.Name, p.Age)
}

func main() {
	jack := Person{"jack", 12} //直接用的struct, 没有引用格式化函数  指针
	tim := &Person{"tim", 32}  // &  引用了格式化函数
	fmt.Println(jack, tim)

}
