package main

import (
	"bufio"
	"fmt"
	"strings"

	"Advent_Of_Code_24/Error"
	"Advent_Of_Code_24/FileReader"
)

type Point struct {
	x, y int
}

type Guard struct {
	pos Point
	dir Point
}

func parseInput(scanner *bufio.Scanner) ([][]byte, Guard) {
	var grid [][]byte
	var guard Guard
	y := 0

	for scanner.Scan() {
		line := []byte(scanner.Text())
		for x, ch := range line {
			if ch == '^' {
				guard = Guard{Point{x, y}, Point{0, -1}}
			}
		}
		grid = append(grid, line)
		y++
	}
	return grid, guard
}

func turnRight(dir Point) Point {
	return Point{-dir.y, dir.x}
}

func isInBounds(p Point, grid [][]byte) bool {
	return p.y >= 0 && p.y < len(grid) && p.x >= 0 && p.x < len(grid[0])
}

func solve(grid [][]byte, guard Guard) int {
	visited := make(map[Point]bool)
	visited[guard.pos] = true

	for {
		nextPos := Point{guard.pos.x + guard.dir.x, guard.pos.y + guard.dir.y}

		if !isInBounds(nextPos, grid) {
			break
		}

		if grid[nextPos.y][nextPos.x] == '#' {
			guard.dir = turnRight(guard.dir)
			continue
		}

		guard.pos = nextPos
		visited[guard.pos] = true
	}

	return len(visited)
}

func solvePart2(grid [][]byte, guard Guard) int {
	startPos := guard.pos
	count := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] != '.' {
				continue
			}

			if x == startPos.x && y == startPos.y {
				continue
			}

			tempGrid := make([][]byte, len(grid))
			for i := range grid {
				tempGrid[i] = make([]byte, len(grid[i]))
				copy(tempGrid[i], grid[i])
			}
			tempGrid[y][x] = '#'

			_, hasLoop := simulateGuard(tempGrid, guard)
			if hasLoop {
				count++
			}
		}
	}

	return count
}

func simulateGuard(grid [][]byte, guard Guard) (map[Point]bool, bool) {
	visited := make(map[Point]bool)
	positions := make(map[Point]bool)
	visited[guard.pos] = true

	state := make(map[string]bool)

	for {
		nextPos := Point{guard.pos.x + guard.dir.x, guard.pos.y + guard.dir.y}

		if !isInBounds(nextPos, grid) {
			return positions, false
		}

		stateStr := fmt.Sprintf("%d,%d,%d,%d", guard.pos.x, guard.pos.y, guard.dir.x, guard.dir.y)
		if state[stateStr] {
			return positions, true
		}
		state[stateStr] = true

		if grid[nextPos.y][nextPos.x] == '#' {
			guard.dir = turnRight(guard.dir)
			continue
		}

		guard.pos = nextPos
		visited[guard.pos] = true
		positions[guard.pos] = true
	}
}

func main() {
	//toString, err := FileReader.ReadFileToString("./Day6/test.txt")
	toString, err := FileReader.ReadFileToString("./Day6/input.txt")
	Error.Check(err)
	scanner := bufio.NewScanner(strings.NewReader(toString))
	grid, guard := parseInput(scanner)

	fmt.Println("Part 1 Result:", solve(grid, guard))
	fmt.Println("Part 2 Result:", solvePart2(grid, guard))
}
