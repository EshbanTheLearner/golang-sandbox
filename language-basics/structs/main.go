package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {

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
}
