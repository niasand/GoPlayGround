package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)
	go inCounter1(1)
	go inCounter1(2)

	wg.Wait()

	fmt.Println("Final counter: ", counter)

}

func inCounter1(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1) //安全的对count加1
		runtime.Gosched()            //当前goroutine从线程退出，并放回到队列
	}
}
