package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type entity struct {
	X int
	Y int
	DirX int
	DirY int
}

func newGuard(x, y int) *entity {
	return &entity{
		x, y, 0, -1,
	}
}

func (guard *entity) move() {
	guard.X += guard.DirX
	guard.Y += guard.DirY
}

func (guard *entity) turn() {
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

func isInBounds(x, y, tileMapWidth, tileMapHeight int) bool {
	if x >= 0 && x < tileMapWidth && y >= 0 && y < tileMapHeight {
		return true
	}
	return false
}

func traceGuardPath(tileMap [][]byte, guard entity) ([][]byte, int) {
	tempTileMap := make([][]byte, len(tileMap))
	for i, row := range tileMap {
		tempTileMap[i] = make([]byte, len(row))
		copy(tempTileMap[i], row)
	}
	movements := 0
	tileMapWidth := len(tempTileMap[0])
	tileMapHeight := len(tempTileMap)
	for isInBounds(guard.X, guard.Y, tileMapWidth, tileMapHeight) {
		tempTileMap[guard.Y][guard.X] = 'X'
		forwardX := guard.X + guard.DirX
		forwardY := guard.Y + guard.DirY
		if isInBounds(forwardX, forwardY, tileMapWidth, tileMapHeight) && tempTileMap[forwardY][forwardX] == '#' {
			guard.turn()
		}
		forwardX = guard.X + guard.DirX
		forwardY = guard.Y + guard.DirY
		if isInBounds(forwardX, forwardY, tileMapWidth, tileMapHeight) && tempTileMap[forwardY][forwardX] == '#' {
			continue
		}
		if isInBounds(forwardX, forwardY, tileMapWidth, tileMapHeight) && tempTileMap[forwardY][forwardX] != 'X' {
			guard.move()
			movements++
		} else {
			guard.move()
		}
	}
	return tempTileMap, movements + 1
}

func isPathLooping(tileMap [][]byte, guard entity) bool {
	tileMapWidth := len(tileMap[0])
	tileMapHeight := len(tileMap)
	for isInBounds(guard.X, guard.Y, tileMapWidth, tileMapHeight) {
		// fmt.Println("X:", guard.X, "Y:", guard.Y)
		tileDirX, tileDirY := decodeDirectionByte(tileMap[guard.Y][guard.X])
		if tileDirX == guard.DirX && tileDirY == guard.DirY {
			return true
		}
		tileValue := tileMap[guard.Y][guard.X]
		if tileValue != '^' && tileValue != '>' && tileValue != 'v' && tileValue != '<' {
			tileMap[guard.Y][guard.X] = getDirectionByte(guard.DirX, guard.DirY)
		}
		forwardX := guard.X + guard.DirX
		forwardY := guard.Y + guard.DirY
		if isInBounds(forwardX, forwardY, tileMapWidth, tileMapHeight) && tileMap[forwardY][forwardX] == '#' {
			guard.turn()
		}
		forwardX = guard.X + guard.DirX
		forwardY = guard.Y + guard.DirY
		if isInBounds(forwardX, forwardY, tileMapWidth, tileMapHeight) && tileMap[forwardY][forwardX] == '#' {
			continue
		}
		guard.move()
	}
	return false
}

func getDirectionByte(dirX, dirY int) byte {
	switch {
	case dirX == 0 && dirY == -1:
		return '^'
	case dirX == 1 && dirY == 0:
		return '>'
	case dirX == 0 && dirY == 1:
		return 'v'
	case dirX == -1 && dirY == 0:
		return '<'
	}
	return '.'
}

func decodeDirectionByte(value byte) (int, int) {
	switch value {
	case '^':
		return 0, -1
	case '>':
		return 1, 0
	case 'v':
		return 0, 1
	case '<':
		return -1, 0
	}
	return 0, 0
}

func countPossibleObstacles(guardPath [][]byte, originalTileMap [][]byte, guard entity) int {
	locations := 0
	pathWidth := len(guardPath[0])
	pathHeight := len(guardPath)
	for i := range pathWidth * pathHeight {
		x := i % pathWidth
		y := i / pathWidth
		if guardPath[y][x] != 'X' {
			continue
		}
		tileMap := make([][]byte, len(originalTileMap))
		for i, row := range originalTileMap {
			tileMap[i] = make([]byte, len(row))
			copy(tileMap[i], row)
		}
		tileMap[y][x] = '#'
		if isPathLooping(tileMap, guard) {
			locations++
		}
	}
	return locations
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
	var tileMap [][]byte
	var guard entity
	y := 0
	for scanner.Scan() {
		tileStrs := strings.Split(scanner.Text(), "")
		newTiles := make([]byte, len(tileStrs))
		for x, tileStr := range tileStrs  {
			tileValue := byte(tileStr[0])
			if tileValue == '^' {
				guard = *newGuard(x, y)
				tileValue = '.'
			}
			newTiles[x] = tileValue
		}
		tileMap = append(tileMap, newTiles)
		y++
	}
	guardPath, movements := traceGuardPath(tileMap, guard)
	fmt.Println("Guard movements:", movements)
	possibleObstacles := countPossibleObstacles(guardPath, tileMap, guard)
	fmt.Println("Possible obstacle locations:", possibleObstacles)
}
