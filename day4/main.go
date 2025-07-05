package main

import (
	"bufio"
	"fmt"
	"os"
)

func findOccurrences(puzzle []string) int {
	word := "XMAS"
	occurrences := 0
	for y := range puzzle {
		for x, char := range puzzle[y] {
			if byte(char) != word[0] {
				continue
			}
			if isMatch(puzzle, x, y, -1, 0, word, 0) {
				occurrences++
			}
			if isMatch(puzzle, x, y, -1, -1, word, 0) {
				occurrences++
			}
			if isMatch(puzzle, x, y, 0, -1, word, 0) {
				occurrences++
			}
			if isMatch(puzzle, x, y, 1, -1, word, 0) {
				occurrences++
			}
			if isMatch(puzzle, x, y, 1, 0, word, 0) {
				occurrences++
			}
			if isMatch(puzzle, x, y, 1, 1, word, 0) {
				occurrences++
			}
			if isMatch(puzzle, x, y, 0, 1, word, 0) {
				occurrences++
			}
			if isMatch(puzzle, x, y, -1, 1, word, 0) {
				occurrences++
			}
		}
	}
	return occurrences
}

func findOccurrencesX(puzzle []string) int {
	occurrences := 0
	for y := range puzzle {
		for x, char := range puzzle[y] {
			if byte(char) == 'A' && isMatchX(puzzle, x, y) {
				occurrences++
			}
		}
	}
	return occurrences
}

func isMatch(puzzle []string, x int, y int, xDir int, yDir int, word string, index int) bool {
	if index >= len(word) {
		return true
	}
	if x < 0 || x >= len(puzzle[0]) || y < 0 || y >= len(puzzle) {
		return false
	}
	if word[index] != puzzle[y][x] {
		return false
	}
	return isMatch(puzzle, x + xDir, y + yDir, xDir, yDir, word, index + 1)
}

func isMatchX(puzzle []string, x, y int) bool {
	if x < 1 || x >= len(puzzle[0]) - 1 || y < 1 || y >= len(puzzle) - 1 {
		return false
	}
	if (puzzle[y - 1][x - 1] == 'M' && puzzle[y + 1][x + 1] == 'S' ||
		puzzle[y - 1][x - 1] == 'S' && puzzle[y + 1][x + 1] == 'M') &&
		(puzzle[y - 1][x + 1] == 'M' && puzzle[y + 1][x - 1] == 'S' ||
		puzzle[y - 1][x + 1] == 'S' && puzzle[y + 1][x - 1] == 'M') {
			return true
	}
	return false
}

func main() {
	if len(os.Args) < 2 {
		panic("Invalid arguments")
	}
	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)
	var puzzle []string
	for scanner.Scan() {
		puzzle = append(puzzle, scanner.Text())
	}
	occurrences := findOccurrences(puzzle)
	fmt.Println("Occurrences:", occurrences)
	occurrencesX := findOccurrencesX(puzzle)
	fmt.Println("Occurrences with X:", occurrencesX)
}
