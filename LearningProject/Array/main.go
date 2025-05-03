package main

import "fmt"

// We can also define the map type instead of how we did it below. These notes should also be on the maps.go file.
type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {

	// Using make. You are defining the len of the array. In this case we are adding 2 elements to our array and then appending to the userNames var.
	// Using make - make([]string, 2, 5) - This is the type, len and cap of the array. So our types are strings, the len of the array made is 2 and the cap -
	// - of the array is 5. So this will allocate the amount of space we need. We can also leave cap blank.
	// This means that we will have those 2 empty elements because nothing has been defined in them yet as we have only appended so far.
	// Because we have a dynamic  array we can't call to a specific element to edit it that way since there is nothing there. It's an empty array waiting to -
	// - be appended to. So you can use userNames := []string{}, userNames[0] = "Tom". That just won't work.
	// While you can still append over the cap, it makes it more efficient if you have an array where you know the exact cap you need. This way Go doesn't -
	// - have to allocate even more space somewhere else that isn't needed. If you append over the cap it will then start allocating more space at another -
	// - address. Make is also available for maps.
	userNames := make([]string, 2, 5)

	userNames[0] = "Tom"
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	// For map's you can only define the intended len of the map. That is all.
	courseRatings := make(floatMap, 3) // map[string]float64 this was in these instead of the floatMap func.

	courseRatings["Go"] = 4.7
	courseRatings["React"] = 4.8
	courseRatings["Angular"] = 4.7

	courseRatings.output()

	//fmt.Println(courseRatings)

	// Using a For Loop to work with arrays, slices and map's.
	// If you don't care about the individual item values and indexes, you can also just write "for range userNames". This may come in handy for the work app.
	// The range keyword essentially exposes two values related to this slice for every iteration of the for loop.
	// You don't have to use index and value that's just what they are you can use whatever you'd like there. For map's it's still the same syntax.

	for index, value := range userNames {
		fmt.Println("Index:", index)
		fmt.Println("value:", value)
	}

	for key, value := range courseRatings {
		fmt.Println("Key:", key)
		fmt.Println("Value:", value)
	}

}
