package main

import "fmt"

type Operate struct {
	Operand_L float64
	Operand_R float64
	Operator  string
	Result    float64
}

func Calc(op Operate) float64 {
	switch op.Operator {
	case "+":
		op.Result = op.Operand_L + op.Operand_R
		break
	case "-":
		op.Result = op.Operand_L - op.Operand_R
		break
	case "*":
		op.Result = op.Operand_L * op.Operand_R
		break
	case "/":
		op.Result = op.Operand_L / op.Operand_R
		break
	default:
		op.Result = 0.0
	}

	return op.Result
}

func main() {
	var op Operate

	fmt.Print("Enter first number : ")
	fmt.Scan(&op.Operand_L)

	fmt.Print("Enter operater : ")
	fmt.Scan(&op.Operator)

	fmt.Print("Enter second number : ")
	fmt.Scan(&op.Operand_R)

	fmt.Printf("Your result : %.2f", Calc(op))
}
