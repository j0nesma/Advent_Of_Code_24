package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"Advent_Of_Code_24/Error"
	"Advent_Of_Code_24/FileReader"
)

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
