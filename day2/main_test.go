package main

import (
	"testing"
)

func TestIsReportSafe(t *testing.T) {
	reports := [][]int {
		{ 7, 6, 4, 2, 1 },
		{ 1, 2, 7, 8, 9 },
		{ 9, 7, 6, 2, 1 },
		{ 1, 3, 2, 4, 5 },
		{ 8, 6, 4, 4, 1 },
		{ 1, 3, 6, 7, 9 },
	}
	expected := []bool {
		true,
		false,
		false,
		false,
		false,
		true,
	}
	for i, levels := range reports {
		if isReportSafe(levels) != expected[i] {
			t.Fatal("Is report safe result did not match expected value")
		}
	}
}

func TestIsReportWithProblemDampener(t *testing.T) {
	reports := [][]int {
		{ 7, 6, 4, 2, 1 },
		{ 1, 2, 7, 8, 9 },
		{ 9, 7, 6, 2, 1 },
		{ 1, 3, 2, 4, 5 },
		{ 8, 6, 4, 4, 1 },
		{ 1, 3, 6, 7, 9 },
	}
	expected := []bool {
		true,
		false,
		false,
		true,
		true,
		true,
	}
	for i, levels := range reports {
		if isReportSafeWithProblemDampener(levels) != expected[i] {
			t.Fatal("Is report safe result did not match expected value")
		}
	}
}
