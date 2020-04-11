//Every stand-alone go program must have a main package
package main

//fmt package - to print out statements
import "fmt"
//time package - to deal with time
import "time"
//os package - to perform OS functionalities
import"os"

//Function of utilizing select block
//Select block is similar to switch but used in case of channels
func Select_func(c chan int, end chan string){

	//Infinite for is used to use all the values in the select block
	//If not used, only first value will be used
	for{

		select{

		case <-c :   fmt.Println("Received")					 
		case <-end : fmt.Println("End")
					//To exit from the program after receiving the end signal
					//If not used, leads to deadlock, as we are using select{} in main()
					 os.Exit(0)
					}
		}

	}

//Entry point for the program
func main() {

	//Making two unbuffered channels
	c := make(chan int)
	end := make(chan string)
	go func(){

		c <- 1
		end <- "END"
	}()
	//When using 2 or more Go Routines, give a sleep time of 1 second after every Go Routine to compensate for concurrency
	time.Sleep(time.Second * 1)
	go Select_func(c, end)
	time.Sleep(time.Second * 1)
	//SELECT - Waits Indefinitely by halting main statement from finishing
	//User has to interrupt to end execution, else the program will be keep on running
	//Returns deadlock - If For loop is finite
	select{}
}