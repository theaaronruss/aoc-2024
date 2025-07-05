package main

import (
	"bufio"
	"fmt"
	"os"
)

func findOccurrences(puzzle []string, word string) int {
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
	occurrences := findOccurrences(puzzle, "XMAS")
	fmt.Println("Occurrences:", occurrences)
}
