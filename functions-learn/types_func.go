package main

import "fmt"

// ...Variadic function ...use when do not know how many arguments are coming
// inside the function it becomes a slice
func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total

}

func typess() {
	// Anonymous Functions
	// Functions without names.
	// Stored in variables

	greet := func(name string) {
		fmt.Println("Hello " + name)
	}

	greet("Anupam")

}
