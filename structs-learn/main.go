package main

import "fmt"

type Address struct {
	City string
}

type User struct {
	Name    string
	Age     int
	Address Address //composition , Inheritance Alternative
}

func cp(u User) { //this will make a copy

}

func ptr(u *User) { //this is change the Original
	u.Age = 30
}

func main() {
	u1 := User{
		Name:    "Anupam",
		Age:     21,
		Address: Address{City: "Jabalpur"},
	}

	u1.Age = 20

	fmt.Println(u1)

	u2 := u1 //this is a Copy

	fmt.Println(u2)

	p2 := &u2
	fmt.Println(p2)
}
