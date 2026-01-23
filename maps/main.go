package main

import "fmt"

type User struct {
	name string
	age  int
}

func main() {
	users := make(map[int]*User)

	users[1] = &User{name: "Anupam", age: 30}
	users[2] = &User{name: "John", age: 25}

	for id, user := range users {
		fmt.Println(id, user.name, user.age)
	}
}
