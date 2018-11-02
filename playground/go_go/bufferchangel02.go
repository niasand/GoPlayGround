package main

import (
	"fmt"
	"sync"
)

var tt sync.WaitGroup

func worker01(name string, ping chan int) {

	defer tt.Done()
	for {
		pong, ok := <-ping
		if !ok {
			fmt.Println("ping channel is closed.", name, ping)
			return
		}

		fmt.Printf("ping: %v,pong: %v\n ", ping, pong)

	}
}

func main() {
	tt.Add(4)

	ping := make(chan int, 10)

	for gr := 1; gr <= 4; gr++ {
		go worker01(fmt.Sprintf("gr number is:%v", gr), ping)
	}

	for post := 1; post <= 10; post++ {
		ping <- post
	}

	close(ping)
	tt.Wait()
}
