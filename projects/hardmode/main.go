package main

/*
Hard Mode Assignment
Create a program that reads the contents of a text file then prints its contents to the terminal.
The file to open should be provided as an argument to the program when it is executed at the terminal. For example, 'go run main.go myfile.txt' should open up the myfile.txt file
To read in the arguments provided to a program, you can reference the variable 'os. Args', which is a slice of type string
To open a file, check out the documentation for the 'Open' function in the 'os' package - https://golang.org/pkg/os/#Open
What interfaces does the 'File' type implement?
If the 'File' type implements the 'Reader interface, you might be able to reuse that io.Copy function!
*/

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//Get the arguments from the first onwards
	args := os.Args[1:]
	//Filename is first arg assuming arg given
	fn := args[0]

	//Open that file for read only as we are not writing this time
	f, err := os.OpenFile(fn, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	//Then iocopy this file into our terminal
	io.Copy(os.Stdout, f)
}
