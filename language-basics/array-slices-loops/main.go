// Arrays, Slices and Loops

package main

import "fmt"

func main() {

	/*
		Slice Declaratoin
	*/

	names := []string{"Bruce Wayne", myName()}

	/*
		Appending to the slice
	*/
	names = append(names, "Tony Stark")
	/*
		For Loop
	*/

	for i, name := range names {
		fmt.Println(i, name)
	}

}

func myName() string {
	return "Eshban Suleman"
}
