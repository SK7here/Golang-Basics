//Every stand-alone go program must have a main package
package main

//fmt package - to print out statements
import "fmt"
//time package - to deal with time
import "time"

//Entry point for the program
func main() {

	//Declaring a channel holds NIL by default
	//By default, channels are unbuffered - can hold only 1 value
	var a chan int //Meant to hold integer values
	fmt.Println(a)

	//Making a channel holds its address by default
	c := make(chan int)
	fmt.Println(c)

	//Sender Go Routine(Anonymous) to send the integer 1
	go func(){

		c <- 1
	}()
	time.Sleep(time.Second * 1)

	//Sniff Go Routine(Anonymous) to receive the integer 1
	//When using 2 or more Go Routines, give a sleep time of 1 second after every Go Routine to compensate for concurrency
	go func(){

		rec := <-c
		fmt.Println(rec)
	}()
	time.Sleep(time.Second * 1)
}
