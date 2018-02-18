package model

import (
	"errors"
)

//PaymentOption interface
type PaymentOption interface {
	ProcessPayment(amount float64) (success bool, err error)
}

//Cash struct
type Cash struct {
	name   string
	amount float64
}

//CreateCash contructor
func CreateCash(name string, amount float64) *Cash {
	return &Cash{
		name:   name,
		amount: amount,
	}
}

//ProcessPayment cash
func (c *Cash) ProcessPayment(amount float64) (success bool, err error) {
	if amount > c.amount {
		err = errors.New("(Cash) Insuficient account amount available :C")
		success = false
		return
	}

	c.amount -= amount
	success = true
	return
}

//CreditCard struct
type CreditCard struct {
	name   string
	amount float64
}

//CreateCreditCard contructor
func CreateCreditCard(name string, amount float64) *CreditCard {
	return &CreditCard{
		name:   name,
		amount: amount,
	}
}

//ProcessPayment cash
func (c *CreditCard) ProcessPayment(amount float64) (success bool, err error) {
	if amount > c.amount {
		err = errors.New("(CreditCard) Insuficient account amount available :C")
		success = false
		return
	}

	c.amount -= amount
	success = true
	return
}
