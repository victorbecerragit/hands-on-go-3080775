// challenges/types-primitive/begin/main.go
package main

import "fmt"

// use type inference to create a package-level variable "val" and assign it a value of "global"
var val string = "Hello"

func main() {
	// create a local variable "val" and assign it an int value
	val := 10

	// print the value of the local variable "val"
  //fmt.Println(val)
	fmt.Printf("%T, local val = %v\n", val, val)

	// print the value of the package-level variable "val"
  printglobalVal()
	// update the package-level variable "val" and print it again
  updateGlobal()
	printglobalVal()
	// print the pointer address of the local variable "val"

	// assign a value directly to the pointer address of the local variable "val" and print the value of the local variable "val"
}

func printglobalVal () {
	fmt.Println("Global val=", val)
}

func updateGlobal() {
 val = "updated global"
}