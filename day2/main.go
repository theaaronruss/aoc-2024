package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func isReportSafe(levels []int) bool {
	isIncreasing := false
	isDecreasing := false
	for i := 1; i < len(levels); i++ {
		level1 := levels[i - 1]
		level2 := levels[i]
		delta := math.Abs(float64(level2 - level1))
		if delta < 1 || delta > 3 {
			return false
		}
		isIncreasing = level1 < level2 || isIncreasing
		isDecreasing = level1 > level2 || isDecreasing
		if isIncreasing && isDecreasing {
			return false
		}
	}
	return true
}

func isReportSafeWithProblemDampener(levels []int) bool {
	if isReportSafe(levels) {
		return true
	}
	for i := range levels {
		tempLevels := make([]int, len(levels))
		copy(tempLevels, levels)
		tempLevels = slices.Delete(tempLevels, i, i + 1)
		if isReportSafe(tempLevels) {
			return true
		}
	}
	return false
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	numSafeReports := 0
	numSafeReportsWithDampener := 0
	for scanner.Scan() {
		levelStrs := strings.Split(scanner.Text(), " ")
		levels := make([]int, len(levelStrs))
		for i, levelStr := range levelStrs {
			level, _ := strconv.Atoi(levelStr)
			levels[i] = level
		}
		if isReportSafe(levels) {
			numSafeReports++
		}
		if isReportSafeWithProblemDampener(levels) {
			numSafeReportsWithDampener++
		}
	}
	fmt.Println("Number of safe reports:", numSafeReports)
	fmt.Println("Number of safe reports with problem dampener:", numSafeReportsWithDampener)
}
