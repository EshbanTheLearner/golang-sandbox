package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	/*
		// Declaration Method 1
		eshban := person{"Eshban", "Suleman"}
		fmt.Println(eshban)

		// Declaration Method 2
		eshban2 := person{firstName: "Eshban", lastName: "Suleman"}
		fmt.Println(eshban2)

		// Declaration Method 3
		var eshban3 person // Zero Values Assigned
		fmt.Println(eshban3)
		fmt.Printf("%+v", eshban3)

		// Updating Struct
		eshban3.firstName = "Eshban"
		eshban3.lastName = "Suleman"
		fmt.Printf("%+v", eshban3)
	*/

	eshban := person{
		firstName: "Eshban",
		lastName:  "Suleman",
		contact: contactInfo{
			email:   "eshban@gmail.com",
			zipCode: 55000,
		},
	}

	fmt.Printf("%+v", eshban)
}
