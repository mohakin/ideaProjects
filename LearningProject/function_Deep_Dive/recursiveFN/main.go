package main

import "fmt"

func main() {
	fact := factorial(5)
	fmt.Println(fact)
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)

	/*	result := 1

		for i := 1; i <= number; i++ {
			result = result * i
		}
		return result*/
}

// Recursion is when a function calls itself. A great example of this is a function that calculates the factorial of a number.
// A factorial if 5 is 5 * 4 * 3 * 2 * 1 => 120
// You have to define an exit condition in every recursive function of else it will cause an infinite loop.
