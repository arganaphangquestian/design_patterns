package main

import (
	"errors"
	"fmt"
)

type IPayment interface {
	Pay() error
}

type Payment struct {
	OrderID     string
	Amount      int64
	OrderUserID string
	IsPaid      bool
	Method      string
}

func (p *Payment) Pay() error {
	fmt.Printf("Order with id %s paid!", p.OrderID)
	return nil
}

func (p *Payment) CheckStatus() bool {
	return p.IsPaid
}

type Tokopedia struct {
	Payment
}

type Midtrans struct {
	Payment
}

type Xendit struct {
	Payment
}

func newTokopedia() IPayment {
	return &Tokopedia{
		Payment: Payment{
			Method: "Tokopedia",
		},
	}
}

func newMidtrans() IPayment {
	return &Midtrans{
		Payment: Payment{
			Method: "Midtrans",
		},
	}
}

func newXendit() IPayment {
	return &Xendit{
		Payment: Payment{
			Method: "Xendit",
		},
	}
}

func newPayment(method string) (IPayment, error) {
	if method == "Tokopedia" {
		return newTokopedia(), nil
	}
	if method == "Midtrans" {
		return newMidtrans(), nil
	}
	if method == "Xendit" {
		return newXendit(), nil
	}
	return nil, errors.New("Wrong payment method")
}

func main() {
	fmt.Println("Hello from Factory Methods")
}
