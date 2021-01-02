package GoExamplesReadOnly

import (
	"fmt"
	"math/rand"
	"sync"
)

var wwgg sync.WaitGroup

func workerban(name string, work chan int) {

	defer wwgg.Done()
	for {
		works, ok := <-work
		if !ok {
			fmt.Println("works is not ok. value is: ", works)
			return
		}

		fmt.Println("works is ok, value is: ", works)

		n := rand.Intn(100)
		fmt.Println("n value is: ", n)
		if n%37 == 0 && n != 0 {
			close(work)
			fmt.Println("关闭works")
			return
		}

		works++
		work <- works
	}
}

func nobuffers() {
	defer wwgg.Wait()

	wwgg.Add(2)
	work := make(chan int)

	go workerban("A", work)
	go workerban("B", work)
	go workerban("C", work)

	work <- 1
}
