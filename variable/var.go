package main

import "fmt"

//var jwtToken := 300000   not allowed here
var jwtToken int = 300000              // allowed here
const LoginToken string = "kkfvjfknjf" //Public

func main() {
	var username string = "Venera"
	fmt.Println(username)
	fmt.Printf("Variable type is: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable type is: %T \n", isLoggedIn)

	var smallVal uint = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable type is: %T \n", smallVal)

	var smallFloat float32 = 255.45554451118562
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable type is: %T \n", anotherVariable)

	//implicit type
	var website = "kmdcksnjfdunhjvubf"
	fmt.Println(website)

	//no var style
	numberOfUser := 30000000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable type is: %T \n", LoginToken)

}
