package main

import "fmt"

func main() {
	//multiple variable declaration
	var x, y int = 3, 4
	fmt.Println(x, y)

	//single variable declaration
	var z int
	fmt.Println(z)

	//type inference
	w := 5
	fmt.Println(w)

	//swaaping no temp variable
	a, b := 3, 4
	a, b = b, a
	fmt.Println(a, b)

	///2 ways of making map
	m := make(map[string]int)
	m["key"] = 10
	fmt.Println(m)

	m2 := map[string]int{"key": 20}
	fmt.Println(m2)

	//delete
	delete(m, "key")
	fmt.Println(m)

	//check if key exists
	value, exists := m["key"]
	fmt.Println(value, exists)

	//length of map
	fmt.Println(len(m))

	//iterating map, order is not guaranteed
	for key, value := range m {
		fmt.Println(key, value)
	}
}
