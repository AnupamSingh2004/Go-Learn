package main

import "fmt"

func ValidateUsername(name string) error {
	if name == "" {
		return fmt.Errorf("name must not be empty")
	}
	if len(name) < 3 {
		return fmt.Errorf("name lenght must be bigger than 3 characters")
	}

	return nil
}

func main() {

	err := ValidateUsername("Anupam")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

}
