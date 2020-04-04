//Every stand-alone go program must have a main package
package main

//fmt package - to print out statements
import "fmt"
//time package - to deal with time
import "time"

//Entry point for the program
func main() {

	//Creating a Buffered channel of size 3, holding any type of values
	c := make(chan interface{}, 3)
	fmt.Println(c)

	//Creating a Structure
	type stud struct{

		Name string
	}
	//Sender Go Routine(Anonymous) to send an integer, two structure objects and a string
	//If the size exceeds 3, for every item after 3rd, once an item is sniffed, nnext item will be sent
	go func(){

		c <- 1
		c <- stud{"Kailash"}
		c <- stud{"SK7"}
		c <- "End of Channel items"

		//"close" is used to indicate end of items in channel when iterating over items in buffered channel
		//If not used, leads to deadlock, as iterator will forever be expecting items in channel
		close(c)
	}()

	//When using 2 or more Go Routines, give a sleep time of 1 second after every Go Routine to compensate for concurrency
	time.Sleep(time.Second * 1)

	//Sniff Go Routine(Anonymous) to receive the channel items
	go func(){

		for i := range c{

			fmt.Println(i)
		}
	}()

	time.Sleep(time.Second * 1)
}