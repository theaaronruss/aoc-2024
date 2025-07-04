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

func parseInstruction(instruction string) int {
	regex := regexp.MustCompile("[0-9]{1,3}")
	matches := regex.FindAllString(instruction, -1)
	num1, _ := strconv.Atoi(matches[0])
	num2, _ := strconv.Atoi(matches[1])
	return num1 * num2
}

func main() {
	file, _ := os.Open("input.txt")
	fileBytes, _ := io.ReadAll(file)
	memory := string(fileBytes)
	instructions := findInstructions(memory)
	sum := 0
	for _, instruction := range instructions {
		sum += parseInstruction(instruction)
	}
	fmt.Println("Result:", sum)
}
