package main

import (
	"fmt"
)

func main() {
	fmt.Println("Functions in golang")
	greeter()
	fmt.Println("Result is:", adder(2, 3))

	total, message := proValues(2, 5, 7, 9)
	fmt.Println("The result is:", total, message)
}

func greeter() {
	fmt.Println("Salam")
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func proValues(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi, it is your result"
}
