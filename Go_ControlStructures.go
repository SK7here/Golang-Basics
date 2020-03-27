//Every stand-alone go program must have a main package
package main

//fmt package - to print out statements
import "fmt"

//Entry point for the program
func main() {

	//"IF" - Finding largest among two numbers
	a := 2
	b := 2

	if a>b {
		fmt.Println("a is greater than b")
	}	else if b>a{
		fmt.Println("b is greater than a")
	}	else{
		fmt.Println("a is equal to b")
	}
	fmt.Println()


	
	//FOR - Printing 10 numbers
	for i:=1 ; i<=10; i++{

		fmt.Println(i)
	}

	//Infinite for
	//for{
	//
	//	fmt.Println("Infinite")
	//}

	//For - using Range(array)
	myarr := []string{"Kailash", "SK7", "Professional in the making"}
	for i, value := range myarr{

		//i holds index ; value holds value in the index
		fmt.Printf("\nIndex : %d, Value : %s", i, value)
	}

	//For - using Range(map)\
	mymap := make(map[string]interface{})
	mymap["Name"] = "Kailash"
	mymap["Roll"] = 70
	for k, v := range mymap{

		fmt.Printf("\nKey : %s, Value : %v", k, v)
	}
	fmt.Println()
	fmt.Println()



	//Continue & Break
		//Print even numbers between 1 to 50
	for i:=1 ; i<100 ; i++{
		if i>50{
			break
		}	else if i%2 != 0{
			continue
		} 	else {
			fmt.Println(i)
		}
	}


	
	//Switch
	day := "Sun"
	switch day{
		case "Mon", "Tue", "Wed", "Thu", "Fri" : 
			fmt.Printf("\n%s is a working day", day)
		case "Sat", "Sun":
			fmt.Printf("\n%s is a holiday", day)
		default:
			fmt.Println("Erroneous data")
	}

}