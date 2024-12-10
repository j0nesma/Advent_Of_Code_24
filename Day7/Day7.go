package main

import (
	"fmt"
	"strconv"
	"strings"

	"Advent_Of_Code_24/Error"
	"Advent_Of_Code_24/FileReader"
)

func main() {
	//toString, err := FileReader.ReadFileToString("./Day7/test.txt")
	toString, err := FileReader.ReadFileToString("./Day7/input.txt")
	arr := strings.Split(toString, "\n")
	Error.Check(err)
	fmt.Printf("Result part1 = %v", solvePart1(arr))
	fmt.Printf("Result part2 = %v\n", solvePart2(arr))
}

func evaluate(nums []int, ops []string) int {
	result := nums[0]
	for i := 0; i < len(ops); i++ {
		if ops[i] == "+" {
			result += nums[i+1]
		} else {
			result *= nums[i+1]
		}
	}
	return result
}

func generateOperators(n int) [][]string {
	if n == 0 {
		return [][]string{{}}
	}

	result := [][]string{}
	subCombinations := generateOperators(n - 1)

	for _, sub := range subCombinations {
		plusCombo := append([]string{}, sub...)
		plusCombo = append(plusCombo, "+")
		mulCombo := append([]string{}, sub...)
		mulCombo = append(mulCombo, "*")
		result = append(result, plusCombo, mulCombo)
	}

	return result
}

func canMakeValue(desiredVal int, nums []int) bool {
	operators := generateOperators(len(nums) - 1)

	for _, ops := range operators {
		if evaluate(nums, ops) == desiredVal {
			return true
		}
	}
	return false
}

func solvePart1(arr []string) int {
	total := 0

	for _, row := range arr {
		if len(row) == 0 {
			continue
		}

		split := strings.Split(row, ":")
		desiredVal, _ := strconv.Atoi(strings.TrimSpace(split[0]))
		values := strings.Fields(strings.TrimSpace(split[1]))

		nums := make([]int, len(values))
		for i, val := range values {
			nums[i], _ = strconv.Atoi(val)
		}

		if canMakeValue(desiredVal, nums) {
			total += desiredVal
		}
	}

	return total
}
