package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

//With type grouping

func adding(a, b int) int {
	return a + b
}

func divide(a, b int) (int, int) {
	return a / b, a % b
}

func calc(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b

	return sum, diff
}

func main() {

	fmt.Println(add(3, 4))

	q, r := divide(10, 5)

	fmt.Println("Result of Division and Mod is :", q, r)

}
