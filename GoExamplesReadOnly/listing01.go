package GoExamplesReadOnly

import (
	"fmt"
	"runtime"
	"sync"
)

func mainfalse() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("开始 goroutine")

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Printf("等待Goroutine结束\n")
	wg.Wait()

	fmt.Println("程序结束")
}
