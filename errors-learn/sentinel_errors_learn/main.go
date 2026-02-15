package main

import (
	"errors"
	"fmt"
)

var ErrorNotFound = errors.New("error not found")

func FindUser(id int) (string, error) {
	if id <= 0 {
		return "", ErrorNotFound
	}
	return "Anupam", nil
}

func main() {
	name, err := FindUser(0)

	if err != nil {
		if errors.Is(err, ErrorNotFound) {
			fmt.Println("User Not Found")
			return
		}
		fmt.Println("Unexpected Error Occured")
	}

	fmt.Println(name)
}
