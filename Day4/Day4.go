package main

import (
	"fmt"
	"strings"

	"Advent_Of_Code_24/Error"
	"Advent_Of_Code_24/FileReader"
)

type Direction struct {
	dx, dy int
}

type WordPosition struct {
	aX, aY int
}

var (
	directions = []Direction{
		{0, 1},   // right
		{1, 0},   // down
		{1, 1},   // right-down
		{-1, 1},  // left-down
		{0, -1},  // left
		{-1, 0},  // up
		{-1, -1}, // left-up
		{1, -1},  // right-up
	}

	diagonals = []Direction{
		{1, 1},   // right-down
		{-1, 1},  // left-down
		{-1, -1}, // left-up
		{1, -1},  // right-up
	}

	XMAS = []rune("XMAS")
	MAS  = []rune("MAS")
)

func main() {
	content, err := FileReader.ReadFileToString("./Day4/input.txt")
	Error.Check(err)

	grid := createGrid(content)

	fmt.Println(findWord(grid, XMAS, directions))
	fmt.Println(findCrossingWords(grid, MAS))
}

func createGrid(content string) [][]rune {
	rows := strings.Split(content, "\n")
	grid := make([][]rune, len(rows))
	for i, row := range rows {
		grid[i] = []rune(row)
	}
	return grid
}

func findWord(grid [][]rune, word []rune, directions []Direction) int {
	rows, cols := len(grid), len(grid[0])
	count := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for _, direction := range directions {
				if found, _, _ := searchInDirection(grid, word, i, j, direction); found {
					count++
				}
			}
		}
	}
	return count
}

func findCrossingWords(grid [][]rune, word []rune) int {
	positions := collectWordPositions(grid, word)
	return countIntersections(positions) / 2
}

func collectWordPositions(grid [][]rune, word []rune) []WordPosition {
	rows, cols := len(grid), len(grid[0])
	var positions []WordPosition

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for _, dir := range diagonals {
				if found, _, _ := searchInDirection(grid, word, i, j, dir); found {
					positions = append(positions, WordPosition{i + dir.dx, j + dir.dy})
				}
			}
		}
	}
	return positions
}

func countIntersections(positions []WordPosition) int {
	crossings := 0
	for i := 0; i < len(positions); i++ {
		count := 0
		for j := 0; j < len(positions); j++ {
			if positions[i] == positions[j] {
				count++
			}
		}
		if count == 2 {
			crossings++
		}
	}
	return crossings
}

func searchInDirection(grid [][]rune, word []rune, startx, starty int, direction Direction) (bool, int, int) {
	rows, cols := len(grid[0]), len(grid)
	wordLen := len(word)

	if !isValidStartPosition(startx, starty, rows, cols) {
		return false, -1, -1
	}

	endx := startx + direction.dx*(wordLen-1)
	endy := starty + direction.dy*(wordLen-1)

	if !isValidEndPosition(endx, endy, cols, rows) {
		return false, -1, -1
	}

	for i := 0; i < wordLen; i++ {
		currX := startx + direction.dx*i
		currY := starty + direction.dy*i
		if grid[currX][currY] != word[i] {
			return false, -1, -1
		}
	}
	return true, endx, endy
}

func isValidStartPosition(x, y, rows, cols int) bool {
	return x >= 0 && x < rows && y >= 0 && y < cols
}

func isValidEndPosition(x, y, cols, rows int) bool {
	return x >= 0 && x < cols && y >= 0 && y < rows
}
