//Every Go package belongs to one package
//main package is the entry point of the Go program

package main

import "fmt" //formatted I/O

func printName() {
	fmt.Println("Anupam")
}

func main() {
	fmt.Println("Hello World")
	printName()

	var name string
	fmt.Scanln(&name)

}

/*
 5️⃣ How Go Code Is Organized

One folder = one package

Files in same folder share package

Executable must have package main

Library code never uses main*/
