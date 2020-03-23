//Every stand-alone go program must have a main package
package main

//fmt package - to print out statements
import "fmt"



//Within the interface, call the methods that need to be implemented by each structure 
type ArithmeticOpr interface {

	Add() float64
	Sub() float64
	Mul() float64
	Div() float64
}


// Struct type `SmallNumbers` - implements the 'ArithmeticOpr' interface by implementing all its methods.
type SmallNumbers struct {

	inp1, inp2 float64
}

func (s SmallNumbers) Add() float64 {

	return s.inp1 + s.inp2
}

func (s SmallNumbers) Sub() float64 {

	return s.inp1 - s.inp2
}

func (s SmallNumbers) Mul() float64 {

	return s.inp1 * s.inp2
}

func (s SmallNumbers) Div() float64 {

	return s.inp1 / s.inp2
}




// Struct type `LargeNumbers` - implements the 'ArithmeticOpr' interface by implementing all its methods.
type LargeNumbers struct {

	inp1, inp2 float64
}

func (l LargeNumbers) Add() float64 {

	return l.inp1 + l.inp2
}

func (l LargeNumbers) Sub() float64 {

	return l.inp1 - l.inp2
}

func (l LargeNumbers) Mul() float64 {

	return l.inp1 * l.inp2
}

func (l LargeNumbers) Div() float64 {

	return l.inp1 / l.inp2
}




//Empty interface can be passed as parameter to function to make it handle any type of arguements
func AllPrinterfunc(i interface{}){

	fmt.Println(i)
}



//Entry point for the program
func main() {

	//Any number OR variety of structures can be allocated to the same interface object
		//For % to work, use Printf instead of Println
    var a ArithmeticOpr

    a = SmallNumbers{3, 2}
	fmt.Printf("Inputs are %v\n\n", a)
    fmt.Printf("Add = %f, Sub = %f, Mul = %f, Div = %f\n\n",  a.Add(), a.Sub(), a.Mul(), a.Div())

    a = LargeNumbers{30, 20}
    fmt.Printf("Inputs are %v\n\n", a)
    fmt.Printf("Add = %f, Sub = %f, Mul = %f, Div = %f\n\n",  a.Add(), a.Sub(), a.Mul(), a.Div())

    AllPrinterfunc(1)
    AllPrinterfunc("Hello")
    AllPrinterfunc(struct {}{})

    //Maps - similar to dictionaries in python
    	//Here type of key is string ; type of value can be anything
    mapper := make(map[string]interface{})
    mapper["Name"] = "Kailash"
    mapper["Age"] = 10
    fmt.Println(mapper)
}