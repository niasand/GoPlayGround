package GoExamplesReadOnly

import (
	"fmt"
	"sync"
)

var vv sync.WaitGroup

func work02(task chan string, gr int) {

	defer vv.Done()

	for {
		tasks, ok := <-task

		if !ok {
			fmt.Println("tasks is is not ok.", tasks)
			return
		}

		fmt.Println("task channel is ok", tasks, gr)
	}
}

func mainfalse() {

	vv.Add(4)
	task := make(chan string, 10)

	for gr := 0; gr < 4; gr++ {
		go work02(task, gr)
	}

	for post := 0; post < 10; post++ {
		task <- fmt.Sprintf("task is :%d and go routine number is: ", post)
	}

	close(task)

	vv.Wait()

}
