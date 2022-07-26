package calculator

//package main

/*func main() {
var operator string
var num1, num2 int
fmt.Print("Please enter First number: ")
fmt.Scanln( & num1)
fmt.Print("Please enter Second number: ")
fmt.Scanln( & num2)
fmt.Print("Please enter Operator (+,-,/,%,*):")
fmt.Scanln( & operator)*/

func calculate(num1, num2 int, operator string) int {
	if num2 == 0 && operator == "/" {
		return -1
	}
	if operator != "+" && operator != "-" && operator != "*" && operator != "/" && operator != "%" {
		return -1
	}
	switch operator {
	case "+":
		return num1 + num2

	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	case "%":
		return num1 % num2
	default:
		return -1
	}

}
