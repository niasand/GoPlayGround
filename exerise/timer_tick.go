package main

import "fmt"
import "time"

func main() {

	for range time.Tick(time.Second) {
		fmt.Println("...")
	}
}
