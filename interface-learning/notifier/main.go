package main

import "fmt"

type Notifier interface {
	Send(message string) error
}

type EmailNotifier struct {
	email string
}

type SMSNotifier struct {
	number string
}

func (e EmailNotifier) Send(message string) error {
	fmt.Println("Sending email to", e.email, "with message", message)
	return nil
}

func (s SMSNotifier) Send(message string) error {
	fmt.Println("Sending SMS to", s.number, "with message", message)
	return nil
}

func Notify(n Notifier, message string) {
	n.Send(message)
}

func main() {
	email := EmailNotifier{email: "anupamsingh@gmail.com"}
	Notify(email, "Hello")
}
