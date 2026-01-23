package main

import "fmt"

func main() {
	marks := make(map[string]int)

	marks["History"] = 90
	marks["Biology"] = 80
	marks["Chemistry"] = 70

	_, ok := marks["Math"]

	if ok {
		fmt.Println("Math is present in the map")
	} else {
		fmt.Println("Math is not present in the map")
	}

	if _, ok := marks["Math"]; ok {
		fmt.Println("Math is present in the map")
	} else {
		fmt.Println("Math is not present in the map")
	}

	fmt.Println(marks)

	arr := make(map[string][]int)

	arr["Anupam"] = []int{1, 2, 3}

	fmt.Println(arr)
	
}
