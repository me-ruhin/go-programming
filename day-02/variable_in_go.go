package main

import "fmt"

func main(){
	// Declaring a variable
	// var <variable_name> <data_type>
	// var <variable_name> <data_type> = <value>
	// <variable_name> := <value>
	// Println is used to print the output in new line
	// Printf is used to print the output in same line
	// %T is used to print the data type of the variable 

	// var name string = " My Name is John Doe"
//	var age int = 25
	// var isMarried bool = false
	// var height float64 = 5.8
	//  name := "ruhin mia"

	// isStudent :=true // inferred data type
	// fmt.Printf("name: %T\n", name)
	// fmt.Println(name)
	// fmt.Println(name)
	// fmt.Printf("age: %s\n", name)
	// fmt.Println(age)
	// fmt.Printf("isMarried: %T\n", isMarried)
	// fmt.Println(isMarried)
	// fmt.Printf("height: %T\n", height)
	// fmt.Println(height)
	// fmt.Println(isStudent)
	// fmt.Printf("isStudent: %T\n", isStudent)
	// fmt.Println(name)

	// variable wihtout value 
	// var country string
	// fmt.Println(country)
	// country = "Bangladesh"
	// fmt.Println(country)

	// variable declaration with multiple variables
	// var a,b,c int = 5, 10, 15
	// fmt.Println(a)
	// fmt.Println(b)	
	// fmt.Println(c)

	// var (
	// 	name string = "John Doe"
	// 	age int = 25
	// 	isMarried bool = false
	// 	height float64 = 5.8)
	// fmt.Println(name)
	// fmt.Println(age)	
	// fmt.Println(isMarried)
	// fmt.Println(height)

	//  const pi float32 = 3.14159
	// pi := 3.14159
	//  fmt.Printf("type of pi: %T\n", pi)
	//  fmt.Println(pi)
	arrayType()
}

func arrayType(){

	// fmt.Println("Array Type")
	// var arr1 = [3]int{}
	var arr2 = [...] string {"a", "b", "c"}
	arr3 := [3]int{1, 2}
	fmt.Println(arr2)
	fmt.Println(len(arr3))
}