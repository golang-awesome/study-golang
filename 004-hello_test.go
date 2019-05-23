package main

import "testing"

func TestNames(t *testing.T) {
	name := getName()
	if name != "World!" {
		t.Error("Unexpected val")
	}
}
