package main

import (
	"fmt"
	"strings"
)

func main() {
	// create two channels. Ping is what we send to, and pong is what comes back.
	ping := make(chan string)
	pong := make(chan string)

	// start a goroutine
	go shout(ping, pong)

	fmt.Println("Type something and press ENTER (enter Q to quit)")
	for {
		fmt.Print("-> ")

		// get user input
		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if userInput == strings.ToLower("q") {
			break
		}

		// send userInput to "ping" channel
		ping <- userInput

		// wait for a response from the pong channel. Again, program
		// blocks (pauses) until it receives something from
		// that channel.
		response := <-pong

		fmt.Println("Response:", response)
	}

	fmt.Println("All done. Closing channels.")
	// close the channels
	close(ping)
	close(pong)
}
