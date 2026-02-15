package main

import (
	"errors"
	"fmt"
)

func readConfig() error {
	return errors.New("simulating error")
}

func loadApp() error {
	err := readConfig()

	if err != nil {
		return fmt.Errorf("load app failed: %w", err)
	}

	return nil
}

func main() {
	err := loadApp()

	if err != nil {
		fmt.Println(err)
		return
	}
}
