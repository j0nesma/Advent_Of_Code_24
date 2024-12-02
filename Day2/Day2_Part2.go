package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"Advent_Of_Code_24/Error"
	"Advent_Of_Code_24/FileReader"
)

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
