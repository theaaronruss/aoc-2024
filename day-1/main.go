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

func getListTotalDistance(list1, list2 []int) int {
	slices.Sort(list1)
	slices.Sort(list2)
	totalDistance := 0
	for i, item1 := range list1 {
		item2 := list2[i]
		distance := math.Abs(float64(item2 - item1))
		totalDistance += int(distance)
	}
	return totalDistance
}

func getListSimilarityScore(list1, list2 []int) int {
	occurrences := make(map[int]int)
	for _, item := range list2 {
		occurrences[item]++
	}
	score := 0
	for _, item := range list1 {
		score += item * occurrences[item]
	}
	return score
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	list1 := make([]int, 1000)
	list2 := make([]int, 1000)
	i := 0
	for scanner.Scan() {
		lineItems := strings.Split(scanner.Text(), "   ")
		list1[i], _ = strconv.Atoi(lineItems[0])
		list2[i], _ = strconv.Atoi(lineItems[1])
		i++
	}
	totalDistance := getListTotalDistance(list1, list2)
	similarityScore := getListSimilarityScore(list1, list2)
	fmt.Println("Total distance between lists:", totalDistance)
	fmt.Println("Similarity score:", similarityScore)
}
