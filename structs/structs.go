package main

import (
	"fmt"
	"example.com/structs/user";
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := user.NewUser(userFirstName, userLastName, userBirthdate)

	appUser.OutputUserDetails()

	if (err != nil) {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@test.com", "password")

	// ... do something awesome with that gathered data!

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
