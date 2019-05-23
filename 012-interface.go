package main

type Pet interface {
	SetName(name string)
	Name() string
	Category() string
}
