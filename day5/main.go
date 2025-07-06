package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func isValidUpdate(update []int, rules map[int][]int) bool {
	if len(update) <= 1 {
		return true
	}
	for i := 1; i < len(update); i++ {
		currPage := update[i]
		rulesForPage := rules[currPage]
		for _, dependency := range rulesForPage {
			if slices.Contains(update[:i], dependency) {
				return false
			}
		}
	}
	return true
}

func parseUpdate(update []string) []int {
	parsedUpdate := make([]int, len(update))
	for i, page := range update {
		pageInt, _ := strconv.Atoi(page)
		parsedUpdate[i] = pageInt
	}
	return parsedUpdate
}

func getMiddlePage(update []int) int {
	midpoint := len(update) / 2
	return update[midpoint]
}

func main() {
	if len(os.Args) < 2 {
		panic("Invalid arguments")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic("Failed to open input file")
	}
	scanner := bufio.NewScanner(file)
	rules := make(map[int][]int)
	for scanner.Scan() && scanner.Text() != "" {
		rule := strings.Split(scanner.Text(), "|")
		page1, _ := strconv.Atoi(rule[0])
		page2, _ := strconv.Atoi(rule[1])
		rules[page1] = append(rules[page1], page2)
	}
	sum := 0
	for scanner.Scan() {
		updateLine := strings.Split(scanner.Text(), ",")
		update := parseUpdate(updateLine)
		if isValidUpdate(update, rules) {
			sum += getMiddlePage(update)
		}
	}
	fmt.Println("Sum of middle pages:", sum)
}
