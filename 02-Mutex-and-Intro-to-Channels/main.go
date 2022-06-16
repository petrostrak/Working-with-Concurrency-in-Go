package main

import (
	"fmt"
	"sync"
)

var (
	wg  sync.WaitGroup
	msg string
)

func updateMessage(s string) {
	defer wg.Done()

	msg = s
}

// go run -race .
func main() {
	msg = "Hello, world!"

	wg.Add(2)
	go updateMessage("Hello, universe!")
	go updateMessage("Hello, cosmos!")

	wg.Wait()

	fmt.Println(msg)
}
