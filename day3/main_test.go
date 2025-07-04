package main

import "testing"

func TestFindInstructions(t *testing.T) {
	matches := findInstructions("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")
	expected := []string {
		"mul(2,4)",
		"mul(5,5)",
		"mul(11,8)",
		"mul(8,5)",
	}
	if len(matches) != len(expected) {
		t.Fatal("Incorrect number of matches")
	}
	for i, match := range matches {
		if match != expected[i] {
			t.Fatal("Incorrect match")
		}
	}
}

func TestParseInstruction(t *testing.T) {
	instructions := []string {
		"mul(2,4)",
		"mul(5,5)",
		"mul(11,8)",
		"mul(8,5)",
	}
	result := parseInstructions(instructions)
	expected := 161
	if result != expected {
		t.Fatalf("Result was %d, not %d", result, expected)
	}
}
