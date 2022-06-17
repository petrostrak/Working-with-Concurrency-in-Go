package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

// go test -race .
func Test_Main(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "$34320.00") {
		t.Error("Wrong balance found")
	}
}
