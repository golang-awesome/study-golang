package main

type SomeStruct struct {
}

func main() {
	s := &SomeStruct{}
	v := SomeStruct{}
	s := &v
	s := new(SomeStruct)

}
