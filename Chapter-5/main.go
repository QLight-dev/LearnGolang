package main

import "fmt"

func main() {
	// array
	// var flavors [5]string = [5]string{"Chocolate", "Vanilla", "Strawberry", "Mango", "Cotton Candy"}
	var flavors = [5]string{"Chocolate", "Vanilla", "Strawberry", "Mango", "Cotton Candy"}
	fmt.Println(flavors, len(flavors))

	// slices (use arrays under the hood)
	var tempuratures = []int{23, 25, 26}
	// place holder
	fmt.Println(tempuratures, len(tempuratures))
}
