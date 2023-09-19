package main

import "fmt"

func main() {
	var (
		x         int
		y         int
		operation string
	)

	fmt.Print("x = ")
	fmt.Scan(&x)
	fmt.Print("y = ")
	fmt.Scan(&y)
	fmt.Print("Choose operation: ")
	fmt.Scan(&operation)

	if operation == "+" {
		println(x + y)
	}

	if operation == "-" {
		println(x - y)
	}

	if operation == "*" {
		println(x * y)
	}

	if operation == "/" {
		println(x / y)
	}
}
