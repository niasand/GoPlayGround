package GoExamplesReadOnly

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	fmt.Println("...")
	return &MyError{
		time.Now(),
		"It does not work",
	}
}

func mainfalse() {
	if err := run(); err == nil {
		fmt.Println(err)
	}
}
