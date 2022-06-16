package main

import (
	"fmt"
	"sync"
)

var (
	wg  sync.WaitGroup
	msg string
)

func printSomething(s string) {
	defer wg.Done()

	fmt.Println(s)
}

func updateMessage(s string) {
	defer wg.Done()

	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {

	words := []string{
		"alpha",
		"beta",
		"gamma",
		"delta",
		"epsilon",
		"zeta",
		"eta",
		"theta",
	}

	for i, x := range words {
		wg.Add(1)
		go printSomething(fmt.Sprintf("%d: %s", i, x))
	}

	printSomething("This is the second thing to be printed")

	// challenge: modify this code so that the calls to updateMessage() on lines
	// 57, 60, and 33 run as goroutines, and implement wait groups so that
	// the program runs properly, and prints out three different messages.
	// Then, write a test for all three functions in this program: updateMessage(),
	// printMessage(), and main().

	msg = "Hello, world!"

	wg.Add(1)
	go updateMessage("Hello, universe!")
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("Hello, cosmos!")
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("Hello, world!")
	wg.Wait()
	printMessage()

}
