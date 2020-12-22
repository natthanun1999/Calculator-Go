package main

import (
	"fmt"
	"testing"
)

func TestCalc(t *testing.T) {
	var op Operate

	assertCorrectMessage := func(t *testing.T, got, want float64) {
		t.Helper()

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("5 + 2", func(t *testing.T) {
		fmt.Println("Testing : " + t.Name())

		op.Operand_L = 5.0
		op.Operator = "+"
		op.Operand_R = 2.0

		got := Calc(op)
		want := 7.0
		assertCorrectMessage(t, got, want)
	})

	t.Run("-8 + 3", func(t *testing.T) {
		fmt.Println("Testing : " + t.Name())

		op.Operand_L = -8.0
		op.Operator = "+"
		op.Operand_R = 3.0

		got := Calc(op)
		want := -5.0
		assertCorrectMessage(t, got, want)
	})

	t.Run("135 - 2.5", func(t *testing.T) {
		fmt.Println("Testing : " + t.Name())

		op.Operand_L = 135.0
		op.Operator = "-"
		op.Operand_R = 2.5

		got := Calc(op)
		want := 132.5
		assertCorrectMessage(t, got, want)
	})

	t.Run("4 * 5", func(t *testing.T) {
		fmt.Println("Testing : " + t.Name())

		op.Operand_L = 4.0
		op.Operator = "*"
		op.Operand_R = 5.0

		got := Calc(op)
		want := 23.0
		assertCorrectMessage(t, got, want)
	})

	t.Run("240 / 12", func(t *testing.T) {
		fmt.Println("Testing : " + t.Name())

		op.Operand_L = 240.0
		op.Operator = "/"
		op.Operand_R = 12.0

		got := Calc(op)
		want := 20.0
		assertCorrectMessage(t, got, want)
	})

	t.Run("4 * 7", func(t *testing.T) {
		fmt.Println("Testing : " + t.Name())

		op.Operand_L = 4.0
		op.Operator = "*"
		op.Operand_R = 7.0

		got := Calc(op)
		want := 28.0
		assertCorrectMessage(t, got, want)
	})
}
