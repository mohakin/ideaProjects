package main

import "fmt"

func main() {
	// Step 1.
	hobbies := [3]string{"Coding", "Video_Games", "Cars"}
	fmt.Println(hobbies)

	// Step 2.
	standalone := hobbies[0]
	fmt.Println(standalone)
	newList := hobbies[1:3]
	fmt.Println(newList)

	// Step 3.
	slicedHobbies := hobbies[0:2]
	//sliceHobbies := hobbies[:2]
	fmt.Println(slicedHobbies)

	// Step 4.
	updatedHobbies := slicedHobbies[1:3]
	fmt.Println(updatedHobbies)
	/*	newUpdatedHobbies := append(hobbies[:1])
		fmt.Println(newUpdatedHobbies)*/

	var courseGoals []string

	courseGoals2 := append(courseGoals, "Get a job.", "Have a better understanding of Go!")
	fmt.Println(courseGoals2)
	courseGoals3 := courseGoals2[0:1]
	courseGoals4 := append(courseGoals3, "Getting laid.", "Better understanding.")
	fmt.Println(courseGoals4)

}
