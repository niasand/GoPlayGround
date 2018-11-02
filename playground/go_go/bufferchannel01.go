package main

import (
	"fmt"
	"golang.org/x/exp/rand"
	"sync"
	"time"
)

var wt sync.WaitGroup

func DoWork(gorouting int, task chan int) {
	defer wt.Done()

	for {
		t, ok := <-task
		if !ok {
			fmt.Println("channel is empty closed", t)
			return
		}

		fmt.Println("channel is not empty,", t, gorouting)
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

	}

}

func main() {
	wt.Add(4)
	x := make(chan int, 10)

	for gr := 1; gr <= 4; gr++ {
		go DoWork(gr, x)
	}

	for post := 1; post <= 10; post++ {
		x <- post
	}

	close(x)

	wt.Wait()
}
