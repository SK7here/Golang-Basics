//Every stand-alone go program must have a main package
package main

//fmt package - to print out statements
import "fmt"

//Defining a user-defined function
func Go_Array_func() {

	//Array declaration & assignment
	var Arr1 [4] int
	Arr1[0] = 1
	Arr1[1] = 2
	Arr1[2] = 3
	Arr1[3] = 4

	Arr2 := [] string{"Hi", "I", "am", "Kailash"}

	fmt.Println("Array 1 is =>", Arr1)
	fmt.Println("Array 2 is =>", Arr2)

	Arr2 = append(Arr2, "alias", "SK7")
	fmt.Println("Array 2 after appending is =>", Arr2)
	
}

//Entry point for the program
func main() {

	//Calling the user defined function from main function
	Go_Array_func()

}

