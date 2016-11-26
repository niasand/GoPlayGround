package main

import (
	"fmt"
	"strings"
)

func square(m int) int {return  m*m}

func prod(m,n int) int {return  m*n}

func add1(r rune) rune {return r + 1}
func main(){
	f := prod
	fmt.Println(f(3,2))
	fmt.Println(strings.Map(add1,"yangzhiwei"))
}

