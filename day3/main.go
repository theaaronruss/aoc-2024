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

func main() {
	file, _ := os.Open("input.txt")
	fileBytes, _ := io.ReadAll(file)
	memory := string(fileBytes)
	instructions := findInstructions(memory)
	result := parseInstructions(instructions)
	fmt.Println("Result:", result)
}
