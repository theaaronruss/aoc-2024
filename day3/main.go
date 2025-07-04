package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func findInstructions(input string) []string {
	regex := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	matches := regex.FindAllString(input, -1)
	if matches == nil {
		return []string{}
	}
	return  matches
}

func findInstructionsWithConditionals(input string) []string {
	regex := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\)`)
	matches := regex.FindAllString(input, -1)
	if matches == nil {
		return []string{}
	}
	return  matches
}

func parseInstructions(instructions []string) int {
	regex := regexp.MustCompile("[0-9]{1,3}")
	sum := 0
	for _, instruction := range instructions {
		matches := regex.FindAllString(instruction, -1)
		num1, _ := strconv.Atoi(matches[0])
		num2, _ := strconv.Atoi(matches[1])
		sum += num1 * num2
	}
	return sum
}

func parseInstructionsWithConditionals(instructions []string) int {
	valueRegex := regexp.MustCompile("[0-9]{1,3}")
	sum := 0
	do := true
	for _, instruction := range instructions {
		if instruction == "do()" {
			do = true
			continue
		} else if instruction == "don't()" {
			do = false
			continue
		}
		if !do {
			continue
		}
		values := valueRegex.FindAllString(instruction, -1)
		num1, _ := strconv.Atoi(values[0])
		num2, _ := strconv.Atoi(values[1])
		sum += num1 * num2
	}
	return sum
}

func main() {
	file, _ := os.Open("input.txt")
	fileBytes, _ := io.ReadAll(file)
	memory := string(fileBytes)
	instructions := findInstructions(memory)
	instructionsWithConditionals := findInstructionsWithConditionals(memory)
	result := parseInstructions(instructions)
	resultWithConditionals := parseInstructionsWithConditionals(instructionsWithConditionals)
	fmt.Println("Result:", result)
	fmt.Println("Result with conditionals:", resultWithConditionals)
}
