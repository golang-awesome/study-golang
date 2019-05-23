// first go demo: hello world
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	name := getName()
	fmt.Println("Hello ", name)
	log.Println("Hello", name)
	log.Println(os.Args)
	log.Println(os.Args[1:])
	fmt.Println(strings.Join(os.Args[0:], " "))

	for inx, arg := range os.Args[1:] {
		fmt.Println(inx, arg)
	}

	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	//for {
	//	fmt.Println("identity")
	//}
}

func getName() string {
	return "World!"
}
