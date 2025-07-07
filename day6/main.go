package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Guard struct {
	X int
	Y int
	DirX int
	DirY int
}

func NewGuard(x, y int) *Guard {
	return &Guard{
		x, y, 0, -1,
	}
}

func (guard *Guard) Move() {
	guard.X += guard.DirX
	guard.Y += guard.DirY
}

func (guard *Guard) Turn() {
	switch {
	case guard.DirX == 0 && guard.DirY == -1:
		guard.DirX = 1
		guard.DirY = 0
	case guard.DirX == 1 && guard.DirY == 0:
		guard.DirX = 0
		guard.DirY = 1
	case guard.DirX == 0 && guard.DirY == 1:
		guard.DirX = -1
		guard.DirY = 0
	case guard.DirX == -1 && guard.DirY == 0:
		guard.DirX = 0
		guard.DirY = -1
	}
}

func loadTileMap(scanner *bufio.Scanner) ([][]byte, *Guard) {
	var tileMap [][]byte
	var guardX int
	var guardY int
	y := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		tileLine := make([]byte, len(line))
		for x, tile := range line {
			tileValue := byte(tile[0])
			if tileValue == '^' {
				guardX = x
				guardY = y
				tileValue = '.'
			}
			tileLine[x] = tileValue
		}
		tileMap = append(tileMap, tileLine)
		y++
	}
	return tileMap, NewGuard(guardX, guardY)
}

func countGuardMovements(tileMap[][]byte, guard *Guard) int {
	movements := 0
	tileMapWidth := len(tileMap[0])
	tileMapHeight := len(tileMap)
	for guard.X >= 0 && guard.X < tileMapWidth && guard.Y >= 0 && guard.Y < tileMapHeight {
		forwardX := guard.X + guard.DirX
		forwardY := guard.Y + guard.DirY
		if forwardX >= 0 && forwardX < tileMapWidth && forwardY >= 0 && forwardY < tileMapHeight && tileMap[forwardY][forwardX] == '#' {
			guard.Turn()
		}
		guard.Move()
		if guard.X < 0 || guard.X >= tileMapWidth || guard.Y < 0 || guard.Y >= tileMapHeight {
			continue
		}
		if tileMap[guard.Y][guard.X] != 'X' {
			movements++
			tileMap[guard.Y][guard.X] = 'X'
		}
	}
	return movements
}

func main() {
	if len(os.Args) != 2 {
		panic("Invalid arguments")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic("Failed to open input file")
	}
	scanner := bufio.NewScanner(file)
	tileMap, guard := loadTileMap(scanner)
	movements := countGuardMovements(tileMap, guard)
	fmt.Println("Number of movements:", movements)
}
