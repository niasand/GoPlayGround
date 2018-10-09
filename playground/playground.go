package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type Duration int64

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42 ")
	}

	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 2, nil
}

func main() {

	var dur Duration
	dur = Duration(1000)
	fmt.Println(dur)

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed", e)
		} else {
			fmt.Println("f1 worked", r)
		}
	}

	_, e := f2(42)

	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

	f("yang")
	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")
	// time.Sleep(time.Second * 10)
	fmt.Scanln()
	fmt.Println("done")
	message := make(chan string)
	m := make(chan string, 2)
	m <- "buffered"
	m <- "channel"
	go func() {
		message <- "ping"

	}()

	msg := <-message
	fmt.Println(msg)

	fmt.Println(<-m)
	fmt.Println(<-m)

	done := make(chan bool, 1)
	go worker(done)
	<-done

	var props *actor.Props = actor.FromFunc(func(c actor.Context) {
		if msg, ok := c.Message().(string); ok {
			fmt.Println(msg) // outputs "Hello World"
		}
	})

	pid := actor.Spawn(props)

	pid.Tell("Hello： World")
	time.Sleep(time.Millisecond * 100)
	pid.GracefulStop() // wait for the actor to stop
}

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}
