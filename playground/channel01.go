package main

import "fmt"

func main() {
	intChan := make(chan int, 50)
	exitChain := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, exitChain)

	for {
		_, ok := <-exitChain
		if !ok {
			break
		}
	}

}

func writeData(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan <- i
		fmt.Println("writeData", i)
	}

	close(intChan)
}

func readData(intChan chan int, exitChain chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("readData 读到数据=%v\n", v)
	}

	exitChain <- true
	close(exitChain)
}
