// every go program starts with package
// package main is used to create an executable program
// import is used to include other packages that this program will use
// fmt is a standard library package that provides formatted I/O functions
// main is the function that will be called when the program is executed
// fmt.Println is used to print the message to the console
// go run hello.go to run the program
// go build hello.go to build the program
// ./hello to run the built program

package main

import (
	"fmt"
) 

func main(){
	fmt.Println("Hello, World!")
}