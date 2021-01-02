package GoExamplesReadOnly

import (
	"fmt"
)

func mainfalse() {
	message := make(chan string)
	go func() { message <- "ping..." }()

	msg := <-message
	fmt.Println(msg)

	buffer()
}

func buffer() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)

	fmt.Println(<-messages)
}
