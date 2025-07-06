package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countGuardMovements(mapTiles [][]string, guardX, guardY int) int {
	movementCount := 0
	dirX := 0
	dirY := -1
	for guardX >= 0 && guardX < len(mapTiles[0]) && guardY >= 0 && guardY < len(mapTiles) {
		if mapTiles[guardY][guardX] != "X" {
			movementCount++
			mapTiles[guardY][guardX] = "X"
		}
		guardX += dirX
		guardY += dirY
		if guardX + dirX < 0 || guardX + dirX >= len(mapTiles[0]) || guardY + dirY < 0 || guardY + dirY >= len(mapTiles) {
			continue
		}
		if mapTiles[guardY + dirY][guardX + dirX] == "#" {
			dirX, dirY = changeDirection(dirX, dirY)
		}
	}
	return movementCount
}

func changeDirection(dirX, dirY int) (int, int) {
	switch {
	case dirX == 0 && dirY == -1:
		dirX = 1
		dirY = 0
	case dirX == 1 && dirY == 0:
		dirX = 0
		dirY = 1
	case dirX == 0 && dirY == 1:
		dirX = -1
		dirY = 0
	case dirX == -1 && dirY == 0:
		dirX = 0
		dirY = -1
	}
	return dirX, dirY
}

func main() {
	if len(os.Args) != 2 {
		panic("Invalid arguments")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic("Failed to read input file")
	}
	scanner := bufio.NewScanner(file)
	var mapTiles [][]string
	var guardX int
	var guardY int
	y := 0
	for scanner.Scan() {
		newTiles := strings.Split(scanner.Text(), "")
		for x, tile := range newTiles {
			if tile == "^" {
				guardX = x
				guardY = y
				newTiles[x] = "."
			}
		}
		mapTiles = append(mapTiles, newTiles)
		y++
	}
	guardMovements := countGuardMovements(mapTiles, guardX, guardY)
	fmt.Println("Guard movements count:", guardMovements)
}
