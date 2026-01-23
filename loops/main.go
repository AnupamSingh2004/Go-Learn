package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	k := []int{1, 2, 3, 4, 5}

	for i, j := range k {
		fmt.Println(i, j)
	}

	day := "Mon"

	switch day {
	case "Mon":
		fmt.Println("Start")
	case "Fri":
		fmt.Println("End")
	default:
		fmt.Println("Mid")
	}

}
