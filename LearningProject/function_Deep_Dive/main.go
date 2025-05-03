package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	sum := sumUp(1, 10, 15, 40, -5)
	anotherSum := sumUp(1, numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

// By using ...int/string/float/[]int with this we're saying it can be called with a list of separate parameter values so long as it's of that type.
// The amount of parameters is now variable which is what makes up a variatic function.
// You can also mix this with other parameters which are defined standalone.
// With this function startingValue will take "1" from the sum slice and 10, 15, 49, -5 will be stored into numbers.

// Normally the function below wants a list of parameters passed through it and not a slice. In order to pass a slice we'll have to go above -
// - and use something similar to how we're passing it through this function. With our anotherSum variable we call the function sumUp -
// as sumUp(1, numbers...). When we do it like this and we pass our slice through with the 3 dots it will actually pull all the elements from the slice -
// and turn that slice into a list of standalone values. The 1 that you see called is for our startingValue that you see called in our sumUp function.
// Sidenote. startingValue is never actually called in this code, so there really isn't a need for it here. With that said you can use something like this -
// - if you wanted to in your own project this is just to show you can pass through more than one argument.
func sumUp(startingValue int, numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val // sum = sum + val
	}
	return sum
}
