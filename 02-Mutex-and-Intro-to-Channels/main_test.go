package main

import (
	"sync"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	msg = "Hello, world!"

	var m sync.Mutex

	wg.Add(1)
	go updateMessage("Goodbye, cruel world!", &m)
	wg.Wait()

	if msg != "Goodbye, cruel world!" {
		t.Error("incorrect value in msg")
	}
}
