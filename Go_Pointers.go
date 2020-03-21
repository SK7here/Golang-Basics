//Every stand-alone go program must have a main package
package main

//fmt package - to print out statements
import "fmt"

//Call by Value method will affect only the dummy variables used within the function
func Go_Pointers_swap_func_CallByValue(m1 int, m2 int ) {

	var temp int
	temp = m1
	m1 = m2
	m2 = temp
	
}

//Call by function method will affect both dummy and actual variables
func Go_Pointers_swap_func_CallByReference(m1 *int, m2 *int ) {

	var temp int
	temp = *m1
	*m1 = *m2
	*m2 = temp
	
}


//Entry point for the program
func main() {

	m1 := 2
	m2 := 3

	fmt.Println("Data of m1 =", m1)
	fmt.Println("Data of m2 =", m2)
		//'&' -> holds address of the variable
	fmt.Println("Addr of m1 =", &m1)
	fmt.Println("Addr of m2 =", &m2)


	fmt.Println("After performing Call By Value")
		//During Call by value, data of variables are sent as parameters
	Go_Pointers_swap_func_CallByValue(m1, m2)
	fmt.Println("After Swapping")
	fmt.Println("m1 =", m1)
	fmt.Println("m2 =", m2)

	fmt.Println("After performing Call By Reference")
		//During Call by reference, addresses of variables are sent as parameters
	Go_Pointers_swap_func_CallByReference(&m1, &m2)
	fmt.Println("After Swapping")
	fmt.Println("m1 =", m1)
	fmt.Println("m2 =", m2)

}