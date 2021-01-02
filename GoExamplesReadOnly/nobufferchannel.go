package GoExamplesReadOnly

import "sync"
import (
	"fmt"
	"math/rand"
	"time"
)

var w sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func play(name string, beat chan int) {
	defer w.Done()

	for {
		ball, ok := <-beat
		fmt.Println("通道正在传输数据")

		if !ok {
			fmt.Printf("Player %s win\n", name)
			return
		}

		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s missed\n", name)

			fmt.Println("通道关闭了")
			close(beat)
			fmt.Println("通道关闭了")
			return
		}

		fmt.Printf("Player %s hit %d\n", name, ball)
		ball++
		beat <- ball
	}
}

func bufferchannel() {
	beat := make(chan int)
	w.Add(2)

	go play("yang1", beat)
	go play("yang2", beat)
	beat <- 1

	w.Wait()
}
