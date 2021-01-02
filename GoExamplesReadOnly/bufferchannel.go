package GoExamplesReadOnly

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4  // 要使用的goroutine的是数量
	taskLoad         = 10 // 要处理的工作的数量
)

var waitg sync.WaitGroup //waitg用来等待程序远程

func init() {
	rand.Seed(time.Now().Unix())
}

func bufferchannels() {
	tasks := make(chan string, taskLoad)
	waitg.Add(numberGoroutines)

	for gr := 1; gr <= numberGoroutines; gr++ {
		go workers(tasks, gr)
	}

	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task channel: %d", post)

	}
	fmt.Println("通道已经关闭了***")
	close(tasks)
	fmt.Println("通道已经关闭了---")
	waitg.Wait()
}

func workers(tasks chan string, worker int) {

	defer waitg.Done() //通知函数已经返回

	for {
		fmt.Println("喂数据给通道")
		task, ok := <-tasks
		fmt.Println("喂数据给通道")

		if !ok {
			fmt.Printf("Worker: %d:shutind done\n", worker)
			return
		}

		fmt.Printf("Worker:%d,started %s\n", worker, task)
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Printf("Worker: %d: completed: %s \n", worker, task)
	}
}
