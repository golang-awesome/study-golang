package main

import (
	"fmt"
	"testing"
)

type HelloInterface interface {
	a() ();
	b() (string, error)
	c(string, int) string
}

// test Alt+Enter
type StandardRegistry struct {
	metrics map[string]interface{}
}

func TestNames(t *testing.T) {
	none := StandardRegistry{
		metrics: make(map[string]interface{}),
	}
	fmt.Println(none)
	name := getName()
	if name != "World!" {
		t.Error("Unexpected val")
		panic("error")
	}
	return
}
