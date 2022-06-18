package main

import "sync"

// The Dining Philosophers problem is well known in computer science circles.
// Five philosophers, numbered from 0 through 4, live in a house where the
// table is laid for them; each philosopher has their own place at the table.
// Their only difficulty – besides those of philosophy – is that the dish
// served is a very difficult kind of spaghetti which has to be eaten with
// two forks. There are two forks next to each plate, so that presents no
// difficulty. As a consequence, however, this means that no two neighbours
// may be eating simultaneously.

var (
	philosophers = []string{"Plato", "Sokrates", "Aristotle", "Niche", "Lao Che"}
	wg           sync.WaitGroup
)

func diningProblem(philosopher string, rightFork, leftFork *sync.Mutex) {
	defer wg.Done()

	// print a message

	// lock both forks

	// print a message

	// unlock the mutexes
}

func main() {
	// print intro

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
}
