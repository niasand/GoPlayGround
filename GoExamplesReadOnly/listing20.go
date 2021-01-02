package GoExamplesReadOnly

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court
		if !ok { //如果不ok，表示通道已经被关闭,如果OK的话，表示通道里面还有数据
			fmt.Printf("Player %s -----\n", name)
			return
		}

		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed ,number is %v \n", name, n)

			close(court) //关闭通道，表示通道里面不会有新的数据了
			return
		}

		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++
		court <- ball
	}
}

func mainfalse() {
	court := make(chan int) //创建一个无缓冲的通道

	wg.Add(2) //计数加2，表示要等待两个goroutine

	go player("jim", court)
	go player("lily", court)

	court <- 2
	wg.Wait()

}
