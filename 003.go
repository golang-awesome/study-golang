package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
)

type SomeStruct struct {
}

func main() {
	s := &SomeStruct{}
	v := SomeStruct{}
	s1 := &v
	s2 := new(SomeStruct)
	fmt.Println(s, s1, s2)
	log.Fatal(3, "fail to download")
	log.Info(5, "info logging")
}
