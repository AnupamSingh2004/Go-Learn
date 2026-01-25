package paymentsystem

import "fmt"

type PaymentGateway interface {
	Pay(amount float64) error
}

type Stripe struct {
	apiKey string
}

type RazorPay struct {
	apiKey string
}

func (s *Stripe) Pay(amount float64) error {
	fmt.Println("Paying", amount, "using Stripe")
	return nil
}

func (r *RazorPay) Pay(amount float64) error {
	fmt.Println("Paying", amount, "using RazorPay")
	return nil
}

func ProcessPayment(p PaymentGateway, amount float64) error {
	return p.Pay(amount)
}

func main() {
	p := Stripe{apiKey: "1234"}
	ProcessPayment(&p, 100)

	r := RazorPay{apiKey: "1234"}
	ProcessPayment(&r, 100)

}
