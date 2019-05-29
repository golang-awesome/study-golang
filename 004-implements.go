package main

import "fmt"

type Closeable interface {
	Close()
}

type File struct{}

func (File) Close() {
	fmt.Println("file...")
}

func main() {
	var s Closeable = File{}
	if _, ok := s.(File); ok {
		fmt.Println("it's a file type...")
	}


}
