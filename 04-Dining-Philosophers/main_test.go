package main

import (
	"testing"
	"time"
)

func Test_Main(t *testing.T) {
	eatTime = 0 * time.Second
	thinkTime = 0 * time.Second
	sleepTime = 0 * time.Second

	for i := 0; i < 100; i++ {
		main()

		if len(orderFinished) != 5 {
			t.Error("wrong number of entries in slice")
		}

		orderFinished = []string{}
	}

}
