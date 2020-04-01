//Every stand-alone go program must have a main package
package main

//fmt package - to print out statements
import "fmt"

//sync package - to use WaitGroup
import "sync"

//Entry point for the program
func main() {

	//Wait Groups - Initializing and setting 2 "done" signals to be expected for ending the program
		//This avoids deadlock
	wg := &sync.WaitGroup{}
	wg.Add(2)

	//Go routines are capable of running concurrently with other functions
	go func(){

		fmt.Println("Hello")
		wg.Done()
	}() //Anonymous/Inline/Lambda Functions

	go func(){

		fmt.Println("World")
		wg.Done()
	}() //Anonymous/Inline/Lambda Functions
	
	//If the required number of (2) "done" signals are received, next statement will be executed
	wg.Wait()
	fmt.Println("MainLine")

}
