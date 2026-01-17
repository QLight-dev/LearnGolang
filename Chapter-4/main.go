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

	// Printf (formatted strings)
	fmt.Printf("i am learning %v and my username is %v", whatImLearing, username)
}
