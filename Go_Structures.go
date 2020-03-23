//Every stand-alone go program must have a main package
package main

//fmt package - to print out statements
import "fmt"

//Defining a structure
type Stud_struct struct{

	Name string
	Roll int
}

//Passing Structure objects as parameters
func (s1 Stud_struct) getName() string {

	return s1.Name	
}

func (s1 Stud_struct) getRoll() int {

	return s1.Roll
}



//Entry point for the program
func main() {

	//Creating a Structure object
	c := Stud_struct{
		Name: "Kailash",
		Roll: 70,
	}

	//For % to work, use Printf instead of Println ; %v holds values and %T holds type of the struct object
	fmt.Printf("Inputs are %v\n\n", c)

	//Calling above defined user-defined fucntions
	fmt.Println("Student name is", c.getName())
	fmt.Println("Student roll is", c.getRoll())

	//Declaring an empty structure
	Empty_Struct := struct {}{}
	fmt.Printf("Type of 'Empty_Struct' is %T\n", Empty_Struct)
	fmt.Printf("Value of 'Empty_Struct' is %v\n", Empty_Struct)


}