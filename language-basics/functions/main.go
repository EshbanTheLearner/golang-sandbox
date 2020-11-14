// Functions

package main

import "fmt"

func main() {

	name := name()
	fmt.Println(name)

	age := age()
	fmt.Println(age)

	height := height()
	fmt.Println(height)
}

func name() string {
	return "Eshban Suleman"
}

func age() int {
	return 24
}

func height() float32 {
	return 5.10
}
