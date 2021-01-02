package GoExamplesReadOnly

import (
	"fmt"
)

func mainfalse() {
	fmt.Println("Counting...")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Done")
}
