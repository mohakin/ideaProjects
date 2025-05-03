package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4}

	double := createTransformer(2)
	triple := createTransformer(3)

	// Anon function can be used here and it will be created in time instead of beforehand.
	// We can not use this anywhere else because it's not a defined function.
	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(transformed)
	fmt.Println(tripled)
	fmt.Println(doubled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// We can call this a function factory becuause its a function that creates functinos
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

// Every anon function is a closure and that means that if you use a variable fom a scope in which the function is created then the value of that variable is locked in.
