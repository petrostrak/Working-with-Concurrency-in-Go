package main

const (
	numberOfPizzas = 10
)

var (
	pizzasMade, pizzasFailed, total int
)

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func main() {
	// seed the random number generator

	// print out a message

	// create a producer

	// run the producer in the background

	// create and run a consumer (customers)

	// print out the ending message
}
