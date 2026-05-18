package main

import "fmt"

func main() {
	x := 1

	switch x {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	}

	//case without the variable passed with switch
	a := 2

	switch {
	case a > 2:
		fmt.Println("Greater than 2")
	case a < 2:
		fmt.Println("a is less than 2")
	}

	//Case of multiple in one
	char := "a"

	switch char {
	case "a", "e", "i", "o", "u":
		fmt.Println("char is a vowel")

	default:
		fmt.Println("Consonant")
	}

}
