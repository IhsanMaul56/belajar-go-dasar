package main

import "fmt"

func main() {
	const firstName string = "Ihsan"
	const lastName = "Maulana"

	fmt.Println(firstName + " " + lastName)

	const (
		country = "Indonesia"
		province = "Jawa Barat"
	)

	fmt.Println("Hallo saya " + firstName + " " + "saya lahir di " + province + ", " + country)
}