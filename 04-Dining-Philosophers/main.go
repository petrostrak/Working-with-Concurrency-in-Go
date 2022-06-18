package main

import (
	"fmt"
	"sync"
	"time"
)

// The Dining Philosophers problem is well known in computer science circles.
// Five philosophers, numbered from 0 through 4, live in a house where the
// table is laid for them; each philosopher has their own place at the table.
// Their only difficulty – besides those of philosophy – is that the dish
// served is a very difficult kind of spaghetti which has to be eaten with
// two forks. There are two forks next to each plate, so that presents no
// difficulty. As a consequence, however, this means that no two neighbours
// may be eating simultaneously.

const (
	HUNGER = 3
)

var (
	philosophers = []string{"Plato", "Sokrates", "Aristotle", "Niche", "Lao Che"}
	wg           sync.WaitGroup
	sleepTime    = 1 * time.Second
	eatTime      = 3 * time.Second
)

func diningProblem(philosopher string, leftFork, rightFork *sync.Mutex) {
	defer wg.Done()

	// print a message
	fmt.Println(philosopher, "is seated.")
	time.Sleep(sleepTime)

	for i := HUNGER; i < 0; i-- {
		fmt.Println(philosopher, "is hungry.")
		time.Sleep(sleepTime)

		// lock both forks
		leftFork.Lock()
		fmt.Printf("\t%s picked up the fork to his left.", philosopher)
		rightFork.Lock()
		fmt.Printf("\t%s picked up the fork to his right.", philosopher)

		// print a message
		fmt.Println(philosopher, "has both forks and is eating.")
		time.Sleep(eatTime)

		// unlock the mutexes
		rightFork.Unlock()
		fmt.Printf("\t%s put down the fork on his right.", philosopher)
		leftFork.Unlock()
		fmt.Printf("\t%s put down the fork on his left.", philosopher)
		time.Sleep(sleepTime)
	}

	// print out done message
	fmt.Println(philosopher, "is satisfied")
	time.Sleep(sleepTime)

	fmt.Println(philosopher, "has left the table")
}

func main() {
	// print intro
	fmt.Println("The Dining Philosophers Problem")
	fmt.Println("-------------------------------")

	// spawn one goroutine for each philosopher
	leftFork := &sync.Mutex{}
	for i := 0; i < len(philosophers); i++ {
		// create a mutex for the right fork
		rightFork := &sync.Mutex{}
		wg.Add(1)
		// call a goroutine
		go diningProblem(philosophers[i], leftFork, rightFork)

		rightFork = leftFork
	}

	wg.Wait()

	fmt.Println("------------------")
	fmt.Println("The table is empty")
}
