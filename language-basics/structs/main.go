package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	eshban := person{"Eshban", "Suleman"}
	fmt.Println(eshban)
	eshban2 := person{firstName: "Eshban", lastName: "Suleman"}
	fmt.Println(eshban2)
}
