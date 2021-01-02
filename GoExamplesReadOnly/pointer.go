package GoExamplesReadOnly

import (
	"fmt"
)

func mainfalse() {
	i, j := 42, 2701
	p := &i //point to i
	pq := &j
	fmt.Println(*pq) // read j through the pointer
	*p = 33
	*p = *pq
	fmt.Println(*p) // read i through the pointer
}
