package GoExamplesReadOnly

import (
	"fmt"
)

func mainfalse() {
	pow := make([]int, 10, 19)
	for i := range pow {
		// fmt.Println(uint(i))
		pow[i] = 1 << uint(i)
	}

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
