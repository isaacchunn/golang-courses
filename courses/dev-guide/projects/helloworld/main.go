// Package == Project == Workspace
// This file belongs to package main.

// There are two different types of packages
// Executable -> generats a file that we can run
// Reusable -> Code used as 'helpers'. Good place to put reusable logic
package main

//Setting a package of main determines the type. The main is used to make executable type packages.

//Shortened form of format for debugging purposes.
//Standard library packages.
import "fmt"

func main() {
	fmt.Println("Hi there!")
}
