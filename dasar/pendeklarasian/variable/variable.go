package main

import "fmt"

func main() {
	var name string

	name = "Ihsan"
	fmt.Println(name)

	name = "Maulana"
	fmt.Println(name)

	var age = 25
	fmt.Println(age)

	var (
		firstName = "Ihsan"
		lastName = "Maulana"
	)

	fullName := firstName + " " + lastName
	fmt.Println(fullName)
}