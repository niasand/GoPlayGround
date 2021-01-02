package GoExamplesReadOnly

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutdown int64
	wg       sync.WaitGroup
)

func doWork(name string) {
	defer wg.Done()

	for {
		fmt.Printf("Doing %s work\n", name)
		time.Sleep(250 * time.Millisecond)

		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutdown %s Down\n", name)
			break
		}
	}

}

func mainfalse() {
	wg.Add(2)
	go doWork("A")
	go doWork("B")

	time.Sleep(1 * time.Second)
	fmt.Println("atomic.LoadInt64(&shutdown)", atomic.LoadInt64(&shutdown))
	fmt.Println("Shutdown now.\n")
	atomic.StoreInt64(&shutdown, 1)
	fmt.Println("atomic.LoadInt64(&shutdown)", atomic.LoadInt64(&shutdown))
	wg.Wait()
}
