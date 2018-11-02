package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func printPrime(prefix string) {
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}

	fmt.Println("Completed", prefix)

}

func main() {
	runtime.GOMAXPROCS(1)
	wg.Add(2)

	fmt.Println("创建goroutines")
	go printPrime("A")
	go printPrime("B")

	fmt.Println("等待goroutine结束")
	wg.Wait()
	fmt.Println("结束程序")

}
