package main

import (
	"testing"
)

func TestGetListTotalDistance(t *testing.T) {
	list1 := []int { 3, 4, 2, 1, 3, 3 }
	list2 := []int { 4, 3, 5, 3, 9, 3 }
	totalDistance := getListTotalDistance(list1, list2)
	expected := 11
	if totalDistance != expected {
		t.Fatalf("List distance is not %d", expected)
	}
}

func TestGetListSimilarityScore(t *testing.T) {
	list1 := []int { 3, 4, 2, 1, 3, 3 }
	list2 := []int { 4, 3, 5, 3, 9, 3 }
	score := getListSimilarityScore(list1, list2)
	expected := 31
	if score != expected {
		t.Fatalf("List similarity score is not %d", expected)
	}
}
