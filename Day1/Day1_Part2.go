package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"Advent_Of_Code_24/Error"
	"Advent_Of_Code_24/FileReader"
)

func part2() {
	toString, err := FileReader.ReadFileToString("./Day1/input.txt")
	//toString, err := FileReader.ReadFileToString("./Day1/test.txt")
	Error.Check(err)
	split := strings.Split(toString, "\r\n")
	var left, right, result []float64

	for _, v := range split[0:] {
		i := strings.Split(v, "  ")
		leftInt, err := strconv.ParseFloat(strings.TrimSpace(i[0]), 64)
		Error.Check(err)
		rightInt, err := strconv.ParseFloat(strings.TrimSpace(i[1]), 64)
		Error.Check(err)
		left = append(left, leftInt)
		right = append(right, rightInt)
	}
	fmt.Printf("left:%v, right:%v\n", left, right)
	sort.Float64s(left)
	sort.Float64s(right)
	fmt.Printf("left:%v, right:%v\n", left, right)

	for i := 0; i < len(left); i++ {
		var count float64 = 0
		for j := 0; j < len(right); j++ {
			if right[j] == left[i] {
				count++
			}
		}
		result = append(result, math.Abs(left[i]*count))
	}
	fmt.Println(result)

	var total float64
	for i := 0; i < len(result); i++ {
		total += result[i]
	}
	fmt.Printf("%f", total)
}
