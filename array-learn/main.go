package main

import "fmt"

func main() {
	var arr [3]int
	arr[0] = 10

	fmt.Println(arr)

	a := []int{1, 2, 3}
	fmt.Println(a)

	//Arrays are value types:

	//Slices

	s := make([]int, 5)

	s[0] = 2

	s = append(s, 10)

	fmt.Println(s)

	g := s[2:3]

	fmt.Println(g)

	newSlice := make([]int, 0)

	newSlice = append(newSlice, 1)
	fmt.Println(len(newSlice), cap(newSlice))
	newSlice = append(newSlice, 2)
	fmt.Println(len(newSlice), cap(newSlice))
	newSlice = append(newSlice, 3)
	fmt.Println(len(newSlice), cap(newSlice))

	fmt.Println(newSlice)
}
