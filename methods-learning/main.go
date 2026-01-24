package main

import "fmt"

type User struct {
	firstName string
	lastName  string
	age       int
}

//“This function belongs to User.”

func (u User) printName() {
	fmt.Println(u.firstName, u.lastName)
}

func (u User) printAge() {
	fmt.Println(u.age)
}

func main() {
	u := User{
		firstName: "Anupam",
		lastName:  "Kumar",
		age:       21,
	}
	u.printName()
	u.printAge()
}
