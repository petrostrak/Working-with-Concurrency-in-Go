package main

import (
	"math/rand"
	"time"

	"github.com/fatih/color"
)

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

func (p *Producer) Close() error {
	ch := make(chan error)

	p.quit <- ch

	return <-ch
}

func pizzaria(pizzaMaker *Producer) {}

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// print out a message
	color.Cyan("The Pizzaria is open for business!")
	color.Cyan("----------------------------------")

	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in the background
	go pizzaria(pizzaJob)

	// create and run a consumer (customers)

	// print out the ending message
}
