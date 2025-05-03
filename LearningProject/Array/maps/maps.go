package main

import "fmt"

func main() {
	/*	websites := []string{"https://google.com", "https://aws.com"}
		fmt.Println(websites[0])*/

	// maps must be written like below this includes none of the whitespace.
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Services"])
	websites["LinkedIn"] = "https://www.linkedin.com" // You can add a key to your map.
	fmt.Println(websites)

	delete(websites, "Google") // You can delete a key from a map.
	fmt.Println(websites)
}

// A map is always dynamic "Thank god".
// For map's you can use anything as a key. This includes ints, strings, floats, structs, interfaces and more I'm sure.
