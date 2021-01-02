package GoExamplesReadOnly

import (
	"fmt"
	"golang.org/x/exp/rand"
	"sync"
)

var wg011 sync.WaitGroup

func kickon(name string, kick chan int) {

	defer wg011.Done()

	for {
		ball, ok := <-kick
		if !ok {
			fmt.Printf("name%s kick", name)
			return
		}
		fmt.Println("ball: %s", ball)

		n := rand.Intn(1000)
		fmt.Println("n 的值", n)
		if n%13 == 0 {
			//如果n等于13 就关闭通道
			fmt.Println("ball: 关闭通道")
			close(kick)

			return
		}
		//如果不是13的倍数，就对ball加1

		ball++
		kick <- ball

	}
}

func nobuffer02() {
	wg011.Add(2)
	kick := make(chan int)
	fmt.Println("ball: 开始go routine")
	go kickon("A", kick)
	go kickon("B", kick)

	kick <- 1

	wg011.Wait()

}
