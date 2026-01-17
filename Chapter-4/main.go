package main

import "fmt"

func main() {
	whatImLearing := "golang"
	username := "QLight-dev"
	// Print
	fmt.Print("hello, ")
	fmt.Print("life! \n")
	fmt.Print("yay \n \n")
	fmt.Println("auto new line")
	fmt.Println("told ya")
	fmt.Println("i am learning", whatImLearing, "and my username is", username)

	// Printf (formatted strings) %_ format specifier
	fmt.Printf("i am learning %v and my username is %v \n", whatImLearing, username)
	// adds quotes around
	fmt.Printf("i am learning %q and my username is %q \n", whatImLearing, username)
	// print type
	fmt.Printf("username is type %T \n", username)
	// print float
	fmt.Printf("this is a float: %f \n", 23.23)
	// limit decimal points
	fmt.Printf("this is a float: %0.2f \n", 23.23)

	// Sprintf (save version of printf)
	var savedStr string = fmt.Sprintf("i am learning %v and my username is %v \n", whatImLearing, username)
	fmt.Printf("saved formatted string: %v", savedStr)
}
