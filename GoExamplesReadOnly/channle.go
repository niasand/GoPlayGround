package GoExamplesReadOnly

import "fmt"

import (
	"fmt"
)

func Count(printhello chan string) {

	printhello <- "hello" //将hello传入channel

}

func mainfalse() {
	channles_list := make([]chan string, 10)
	for i := 0; i < 10; i++ {
		channles_list[i] = make(chan string)
		go Count(channles_list[i])
	}

	for _, printhello := range channles_list {
		// <-ch //通过这一句从10个channel中依次读取数据
		fmt.Println("print hello channle", <-printhello)
	}
}
