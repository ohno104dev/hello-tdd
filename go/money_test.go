package main

import (
	"testing"
)

func TestMultiplicationInDollars(t *testing.T) {
	fiver := Money{
		amount:   5,
		currency: "USD",
	}
	actualResult := fiver.Times(2)
	expectedResult := Money{
		amount:   10,
		currency: "USD",
	}
	assertEqual(t, actualResult, expectedResult)
}

func TestMultiplicationInEuros(t *testing.T) {
	tenEuros := Money{
		amount:   10,
		currency: "EUR",
	}
	actualResult := tenEuros.Times(2)
	expectedResult := Money{
		amount:   20,
		currency: "EUR",
	}
	assertEqual(t, actualResult, expectedResult)
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
	actualResult := originalMoney.Divide(4)
	expectedResult := Money{
		amount:   1000.5,
		currency: "KRW",
	}
	assertEqual(t, actualResult, expectedResult)
}

func (m Money) Divide(divisor int) Money {
	return Money{
		amount:   m.amount / float64(divisor),
		currency: m.currency,
	}
}

func assertEqual(t *testing.T, expected Money, actual Money) {
	if expected != actual {
		t.Errorf("Expected %+v Got %+v", expected, actual)
	}
}
