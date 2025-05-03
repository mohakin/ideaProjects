package main

import (
	"fmt"
	"os/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter birthdate (MM/DD/YYYY): ")

	var appUser *user.User

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println("Error! You must fill out all fields!")
		return
	}

	// Do something awesome wth that gathered data!

  appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}


func getUserData(promtText string) string {
	fmt.Println(promtText)
	var value string
	fmt.Scanln(&value)
	return value
}

// ----------------------------------------------------------------
// Structs and Custom Types ---------------------------------------
// ----------------------------------------------------------------
// Struct is a value type that
//
//
//
//
//
//
//
//
//
