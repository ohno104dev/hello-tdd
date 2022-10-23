package main

import (
	"testing"
)

func TestMultiplicationInDollars(t *testing.T) {
	fiver := Money{
		amount:   5,
		currency: "USD",
	}
	tenner := fiver.Times(2)
	if tenner.amount != 10 {
		t.Errorf("Expected 10, got: [%+v]", tenner.amount)
	}
	if tenner.currency != "USD" {
		t.Errorf("Expected USD, got: [%s]", tenner.currency)
	}
}

func TestMultiplicationInEuros(t *testing.T) {
	tenEuros := Money{
		amount:   10,
		currency: "EUR",
	}
	twentyEuros := tenEuros.Times(2)
	if twentyEuros.amount != 20 {
		t.Errorf("Expected 20, got: [%+v]", twentyEuros.amount)
	}

	if twentyEuros.currency != "EUR" {
		t.Errorf("Expected EUR, got: [%s]", twentyEuros.currency)
	}
}

type Money struct {
	amount   float64
	currency string
}

func (m Money) Times(multiplier int) Money {
	return Money{
		amount:   m.amount * float64(multiplier),
		currency: m.currency,
	}
}

func TestDivision(t *testing.T) {
	originalMoney := Money{
		amount:   4002,
		currency: "KRW",
	}
	actualMoneyAfterDivision := originalMoney.Divide(4)
	expectedMoneyAfterDivision := Money{
		amount:   1000.5,
		currency: "KRW",
	}
	if expectedMoneyAfterDivision != actualMoneyAfterDivision {
		t.Errorf("Expected %+v Got %+v", expectedMoneyAfterDivision, actualMoneyAfterDivision)
	}
}

func (m Money) Divide(divisor int) Money {
	return Money{
		amount:   m.amount / float64(divisor),
		currency: m.currency,
	}
}
