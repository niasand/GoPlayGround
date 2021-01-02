package GoExamplesReadOnly

import (
	"fmt"
	"sync"
)

var (
	myMap = make(map[int]int, 10)
	lock  sync.Mutex
)

func lockexmple() {
	for i := 1; i < 10; i++ {
		go test(1)
	}

	// time.Sleep(time.Second * 1)
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()
}

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}
