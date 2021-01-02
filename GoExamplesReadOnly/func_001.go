package GoExamplesReadOnly

import (
	"fmt"
	"strings"

)

func square(m int) int {return  m*m}

func prod(m,n int) int {return  m*n}

func add1(r rune) rune {return r + 1}
func mainfalse(){
	i, j := 0,1
	i, j = j, i
	fmt.Println(i,j)

	f := prod
	fmt.Println(f(3,2))
	fmt.Println(strings.Map(add1,"yangzhiwei"))
}

