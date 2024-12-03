package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"Advent_Of_Code_24/Error"
	"Advent_Of_Code_24/FileReader"
)

func main() {
	fmt.Println("Part 1")
	fmt.Println("")
	part1()
	fmt.Println("")
	fmt.Println("========================================")
	fmt.Println("Part 2")
	fmt.Println("")
	part2()
}

func part1() {
	toString, err := FileReader.ReadFileToString("./Day2/input.txt")
	//toString, err := FileReader.ReadFileToString("./Day2/test.txt")
	Error.Check(err)
	rows := strings.Split(toString, "\r\n")
	var safe = 0
	var isDecending = false
	for _, v := range rows {
		splitRow := strings.Split(v, " ")
		var safeCheck = 0
		for i, val := range splitRow {
			if i == len(splitRow)-1 {
				if safeCheck == len(splitRow)-1 {
					safe++
				}
				break
			}
			nextVal, err := strconv.ParseFloat(splitRow[i+1], 64)
			Error.Check(err)
			floatVal, err := strconv.ParseFloat(val, 64)
			Error.Check(err)
			if i == 0 {
				isDecending = math.Signbit(floatVal - nextVal)
			} else {
				if isDecending != math.Signbit(floatVal-nextVal) {
					break
				}
			}
			if math.Abs(floatVal-nextVal) > 3 || math.Abs(floatVal-nextVal) < 1 {
				break
			}
			safeCheck++
		}
	}
	fmt.Printf("Number of safe rows are: %v", safe)
}

func part2() {
	//toString, err := FileReader.ReadFileToString("./Day2/test.txt")
	toString, err := FileReader.ReadFileToString("./Day2/input.txt")
	Error.Check(err)
	rows := strings.Split(toString, "\r\n")
	var safeAfterRemoval = 0

	for _, v := range rows {
		splitRow := strings.Split(v, " ")
		// Try removing each number one at a time
		for skipIndex := range splitRow {
			var isDecending = false
			var safeCheck = 0
			var lastCheckedIndex = -1

			for i := range splitRow {
				if i == skipIndex {
					continue // Skip this number
				}

				if lastCheckedIndex == -1 {
					lastCheckedIndex = i
					continue
				}

				currentVal, _ := strconv.ParseFloat(splitRow[i], 64)
				lastVal, _ := strconv.ParseFloat(splitRow[lastCheckedIndex], 64)

				if safeCheck == 0 {
					isDecending = math.Signbit(lastVal - currentVal)
				} else {
					if isDecending != math.Signbit(lastVal-currentVal) {
						break
					}
				}

				if math.Abs(lastVal-currentVal) > 3 || math.Abs(lastVal-currentVal) < 1 {
					break
				}

				safeCheck++
				lastCheckedIndex = i

				if safeCheck == len(splitRow)-2 { // -2 because we removed one number
					safeAfterRemoval++
					goto nextRow
				}
			}
		}
	nextRow:
	}
	fmt.Printf("Number of rows that can be made safe by removing one number: %v", safeAfterRemoval)
}
