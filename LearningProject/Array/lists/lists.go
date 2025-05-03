package main

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99
	prices = prices[1:]

	updatedPrices := append(prices, 5.99) // append will add an item to a slice and the underlying array. It will make a new array with the proper size if it isn't already that size.
	fmt.Println(updatedPrices, prices)
}

/*func main() {

var productNames [4]string = [4]string{"A Book"}
prices := [4]float64{10.99, 9.99, 45.99, 20.0}
fmt.Println(prices)
productNames[2] = "A Carpet"
fmt.Println(productNames)
fmt.Println(prices[2])

/*Sliced using [X:X] counts from the first number and goes up to the last number but does not include the last number*/
// You can also do [:X] which will grab the first number of the array all the way up to the last number but not including the last.
// You can also do [X:] which will grab everything from the first number in that you typed in to the end of the array.
// You can't use negative numbers, and you can not go past the end of the array.

/*	featuredPrices := prices[1:]
	featuredPrices[0] = 199.99
	highlightedPrices := featuredPrices[:1]
	fmt.Println(highlightedPrices)
	fmt.Println(prices)
	fmt.Println(len(featuredPrices), cap(featuredPrices)) // len gives us the number of items in an array or slice.
	// cap gives us the end of an array - still confused about this, I will need to work on this more later.
	// Cap will take what you have and more so long as it's added to the array/slice. If you remove something from the array/slice it will not grab it, specially if it's at the beginning.
	highlightedPrices = highlightedPrices[:3]
	fmt.Println(highlightedPrices)
	fmt.Println(len(featuredPrices), cap(featuredPrices))*/
//}
//
