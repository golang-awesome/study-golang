package main

import "fmt"

func main() {
	ints := []int{1, 2, 3, 4,}
	for _, v := range multiply(add(multiply(ints, 2), 1), 10) {
		fmt.Println(v)
	}
	fmt.Println()
	for _, v := range ints {
		fmt.Println(multiplyOne(addOne(multiplyOne(v, 2), 1), 10))
	}
}

var multiply = func(values []int, multiplier int) []int {
	multipliedValues := make([]int, len(values))
	for i, v := range values {
		multipliedValues[i] = v * multiplier
	}
	return multipliedValues
}

var add = func(values []int, additive int) []int {
	addedValues := make([]int, len(values))
	for i, v := range values {
		addedValues[i] = v + additive
	}
	return addedValues
}

var multiplyOne = func(value, multiplier int) int {
	return value * multiplier
}

var addOne = func(value, additive int) int {
	return value + additive
}
