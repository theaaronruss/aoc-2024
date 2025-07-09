package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isValidEquation(targetValue int64, result int64, operands []int) bool {
	if len(operands) == 0 {
		return result == targetValue
	}
	isAddValid := false
	isMultiplyValid := false
	for _, op := range [2]byte {'+', '*'} {
		value := result
		if op == '+' {
			value  += int64(operands[0])
			isAddValid = isValidEquation(targetValue, value, operands[1:])
		} else if op == '*' {
			value *= int64(operands[0])
			isMultiplyValid = isValidEquation(targetValue, value, operands[1:])
		}
	}
	return isAddValid || isMultiplyValid
}

func main() {
	if len(os.Args) != 2 {
		panic("Invalid arguments")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic("Unable to open input file")
	}
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		equations := strings.Split(scanner.Text(), ":")
		result, _ := strconv.Atoi(equations[0])
		operandStrs := strings.Split(equations[1], " ")[1:]
		operands := make([]int, len(operandStrs))
		for i, operandStr := range operandStrs {
			operands[i], _ = strconv.Atoi(operandStr)
		}
		if isValidEquation(int64(result), int64(operands[0]), operands[1:]) {
			sum += result
		}
	}
	fmt.Println("Sum:", sum)
}
