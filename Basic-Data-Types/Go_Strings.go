//Every stand-alone go program must have a main package
package main

//fmt package - to print out statements
//strings package - to use string functions
import "fmt"
import "strings"

//Entry point for the program
func main() {

	//Finding whether a string is a substring of another string
	var s1 string
	var s2 string
	s1 = "Kailash is my name"
	s2 = "Nickname of Kailash is SK7"
	s3 := "Kailash"
	fmt.Println("Statement 1 is " + s1)
	fmt.Println("Statement 2 is " + s2)
	fmt.Println("Statement 3 is " + s3)
	fmt.Println("Statement1 contains Statement2?")
	fmt.Println(strings.Contains(s1, s2))
	fmt.Println("Statement1 contains Statement3?")
	fmt.Println(strings.Contains(s1, s3))


	//Replace characters in strings
	fmt.Println("Statement2 after ReplaceAll operation")
	fmt.Println(strings.ReplaceAll(s2, "Nickname", "PetName"))


	//Split string based on a delimiter
	Splitted_s2 := strings.Split(s2, " ")
	fmt.Println("Words in the Statement2 are")
	fmt.Println(Splitted_s2[0])
	fmt.Println(Splitted_s2[1])
	fmt.Println(Splitted_s2[2])
	fmt.Println(Splitted_s2[3])
	fmt.Println(Splitted_s2[4])

}
