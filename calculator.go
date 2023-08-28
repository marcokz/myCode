package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	result, err := calculate(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result)
}

func calculate(args []string) (int, error) {
	if len(args) != 3 {
		return 0, errors.New("invalid number of arguments")
	}

	x, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, errors.New("first value is not an integer")
	}

	y, err := strconv.Atoi(args[2])
	if err != nil {
		return 0, errors.New("second value is not an integer")
	}

	operation := args[1]

	if operation == "+" {
		return add(x, y)
	}

	if operation == "-" {
		return substract(x, y)
	}

	if operation == "*" {
		return multiply(x, y)
	}

	if operation == "/" {
		return divide(x, y)
	}

	return 0, errors.New("invalid operation")
}

func add(x, y int) (int, error) {
	return x + y, nil
}

func substract(x, y int) (int, error) {
	return x - y, nil
}

func multiply(x, y int) (int, error) {
	return x * y, nil
}

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("divide by zero")
	}
	return x / y, nil
}
