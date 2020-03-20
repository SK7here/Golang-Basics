//Every stand-alone go program must have a main package
package main

//fmt package - to print out statements
import "fmt"

//Entry point for the program
func main() {

	//Actual way to declare variables
	var m1 int
	var m2 int
	m1 = 1
	m2 = 2
	fmt.Println("m1 is", m1)
	fmt.Println("m2 is", m2)
	fmt.Println("Sum of m1 and m2 is", m1 + m2)


	//If multiple variables are of same datatypes
	var(
		m3 = 3
		m4 = 4
		)
	fmt.Println("m3 is", m3)
	fmt.Println("m4 is", m4)
	fmt.Println("Sum of m3 and m4 is", m3 + m4)


	//Typecasting
	var m5 int32
	var m6 int64
	m5 = 5
	m6 = 6
	fmt.Println("m5 is", m5)
	fmt.Println("m6 is", m6)
	fmt.Println("Sum of m5 and m6 is", int64(m5) + m6)


	// '=' is for assignment
	// ':=' is declaration + assignmnet
	//By default, integer variables are assigned with ZERO
	m7 := 7
	m8 := 8
	fmt.Println("m7 is", m7)
	fmt.Println("m8 is", m8)
	fmt.Println("Sum of m7 and m8 is", m7 + m8)

}