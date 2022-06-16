package main

import (
	"fmt"
	"sync"
)

var (
	wg  sync.WaitGroup
	msg string
)

func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()

	m.Lock()
	msg = s
	m.Unlock()
}

// go run -race .
func main() {
	msg = "Hello, world!"

	var m sync.Mutex

	wg.Add(2)
	go updateMessage("Hello, universe!", &m)
	go updateMessage("Hello, cosmos!", &m)

	wg.Wait()

	fmt.Println(msg)
}
