package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int //counter是所有goroutine都要增加他的值的变量
	wg      sync.WaitGroup
)

func inCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		value := counter  //捕获counter的值
		runtime.Gosched() //将当前的goroutine从线程退出，并放回到队列
		value++
		counter = value //将该值保存到counter

	}
}

func main() {
	wg.Add(2)
	go inCounter(1)
	go inCounter(2)
	wg.Wait()
	fmt.Println("Final Counter:", counter)

}

//go build -race listing09.go 来检测是否存在race状态
