package main

import "fmt"
func main() {
	// strings
	var stringOne string = "this is a string"
	var stringTwo = "this is another string"
	var stringThree string
	fmt.Println(stringOne, stringTwo, stringThree)

	stringOne = "this is an updated string"
	stringTwo = "this is another updated string"
	stringThree = "declared this now"
	fmt.Println(stringOne, stringTwo, stringThree)

	stringFour := "didnt write var here"
	fmt.Println(stringFour)

	// ints
	var intOne int = 12
	var intTwo = 34
	var intThree int // 0
	intFour := 56
	fmt.Println(intOne, intTwo, intThree, intFour)

	// bits & memory
	var bitOne int8 = 127 // -128 to 127
	var bitTwo int16 = 32767 // -32768 to 32767
	var bitThree uint8 = 255 // no negative, 0 - 255
	fmt.Println(bitOne, bitTwo, bitThree)

	// floats
	var floatOne float32 = 34.87
	var floatTwo float64 = 739458734987345987345987.23232
	floatThree := 9384.93
	fmt.Println(floatOne, floatTwo, floatThree)
}
