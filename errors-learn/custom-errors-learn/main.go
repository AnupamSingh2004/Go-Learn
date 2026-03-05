package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Field  string
	Reason string
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("validation failed on field '%s' : %s", v.Field, v.Reason)
}

func ValidateAge(age int) error {
	if age < 18 {
		return ValidationError{
			Field:  "age",
			Reason: "must be 18",
		}
	}

	return nil

}

func main() {
	err := ValidateAge(15)

	if err != nil {
		var ve ValidationError
		if errors.As(err, &ve) {
			fmt.Println("Field: ", ve.Field)
			fmt.Println("Reason: ", ve.Reason)
			return
		}

		fmt.Println("Other Error", err)
	}
}
