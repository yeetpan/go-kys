package main

import "fmt"

const LoginToken string = "nfkjlahjkbgdkjajklfhl0" //Public

// jwtToken:=3000 is not allowed use var outside of methods.
func main() {
	var c int = 12
	var username string = "anyways"
	fmt.Println(c, username)
	// %T-> for finding out the type of the variable.
	fmt.Printf("Variable is of the type  %T \n", c)

	var isLoggedin bool = false
	fmt.Print(isLoggedin)
	// limited upto certain values.
	var smallVal uint8 = 2
	fmt.Println(smallVal)

	var smallFloat float64 = 255.4555
	fmt.Println(smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Printf("Variable is of type : %T \n", anotherVariable)

	// implicit type

	var website = "lco.in"
	fmt.Println(website)

	// no var style
	numberOfUser := 30000
	fmt.Println(numberOfUser)
}
