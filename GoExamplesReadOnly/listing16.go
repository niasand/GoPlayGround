package GoExamplesReadOnly

// 互斥锁。互斥锁在代码上创建一个临界区，保证同一时间只有一个goroutine可以执行这个临界区的代码。
// 直到锁被释放。不然在临界区间的那段代码块被不断的被同一个goroutine所执行
import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter01 int

	wg01 sync.WaitGroup

	mutex sync.Mutex
)

func inCounter01(id int) {
	defer wg01.Done()

	for count := 0; count < 2; count++ {
		//同一时刻只允许一个goroutine进入
		//这个临界区
		mutex.Lock()
		{
			value := counter01
			runtime.Gosched() //当前goroutine从线程退出，并放回到队列
			value++
			counter01 = value
		}
		mutex.Unlock() //释放锁，允许其他正在等待的goroutine
		//进入临界区
	}
}

func mainfalse() {
	wg01.Add(2)
	go inCounter01(1)
	go inCounter01(2)

	wg01.Wait()
	fmt.Printf("Final Counter; %d\n", counter01)
}
