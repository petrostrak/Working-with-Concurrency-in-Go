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

func diningProblem() {
	defer wg.Done()
}

func main() {
	// print intro

	// spawn one goroutine for each philosopher
	for i := 0; i < len(philosophers); i++ {
		wg.Add(1)
		// call a goroutine
		go diningProblem()
	}

	wg.Wait()
}
