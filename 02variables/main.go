package main
import "fmt"

// Theis aint allowed outsite a method jwtToken := 300000
const LoginToken string = "Gibrsivdwwi" // This is a public variable; it can be accessed by other fiels as it is exported

func main () {
	var username string = "hitesh"
	fmt.Println(username)
	fmt.Printf("Variabe is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variabe is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variabe is of type: %T \n", smallVal)

	var smallFloat float32 = 255.5567890090
	fmt.Println(smallFloat)
	fmt.Printf("Variabe is of type: %T \n", smallFloat)

	// Default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	//Implicit type
	var website = "www.com"
	fmt.Println((website))

	// no var style
	numberOfUser := 300000
	fmt.Println((numberOfUser))

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}