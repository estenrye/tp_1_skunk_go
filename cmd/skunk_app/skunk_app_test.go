package main

import (
	"testing"
)

func TestMainGreeting(t *testing.T) {
	var expected = "Hello World!"
	var actual = getGreeting()
	if actual != expected {
		t.Errorf("getGreeting() returned incorrect greeting.")
	}
}
