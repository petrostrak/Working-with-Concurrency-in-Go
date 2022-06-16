package main

import (
	"fmt"
	"sync"
)

func printSomething(s string) {
	defer wg.Done()

	fmt.Println(s)
}

var wg sync.WaitGroup

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

	wg.Wait()
}
