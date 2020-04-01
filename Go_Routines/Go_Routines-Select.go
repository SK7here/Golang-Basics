//Every stand-alone go program must have a main package
package main

//fmt package - to print out statements
import "fmt"
//time package - to deal with time
import "time"


func Heavy(){

	//Infinite For loop
	for{
			//After every second, prints "Heavy"
			time.Sleep(time.Second * 1)
			fmt.Println("Heavy")
	}
}


func SuperHeavy(){

	//Infinite For loop
	for{
			//After every 2 seconds, prints "SuperHeavy"
			time.Sleep(time.Second * 2)
			fmt.Println("SuperHeavy")
	}
}


//Entry point for the program
func main() {

	//Go routines are capable of running concurrently with other functions
	//If Go routines are called in main function, other statements in main function will first be executed.
	go Heavy()
	go SuperHeavy()
	fmt.Println("MainLine")
	//Wait Indefinitely
	//Retruns deadlock - If For loop is finite
	select{}
}