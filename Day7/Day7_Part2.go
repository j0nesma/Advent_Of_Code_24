package main

import (
	"strconv"
	"strings"
)

func concatenate(a, b int) int {
	bStr := strconv.Itoa(b)
	aStr := strconv.Itoa(a)
	result, _ := strconv.Atoi(aStr + bStr)
	return result
}

func evalP2(nums []int, ops []string) int {
	result := nums[0]
	for i := 0; i < len(ops); i++ {
		switch ops[i] {
		case "+":
			result += nums[i+1]
		case "*":
			result *= nums[i+1]
		case "||":
			result = concatenate(result, nums[i+1])
		}
	}
	return result
}

func genOperators(n int) [][]string {
	if n == 0 {
		return [][]string{{}}
	}

	result := [][]string{}
	subCombinations := genOperators(n - 1)

	for _, sub := range subCombinations {
		plusCombo := append([]string{}, sub...)
		plusCombo = append(plusCombo, "+")
		mulCombo := append([]string{}, sub...)
		mulCombo = append(mulCombo, "*")
		concatCombo := append([]string{}, sub...)
		concatCombo = append(concatCombo, "||")
		result = append(result, plusCombo, mulCombo, concatCombo)
	}

	return result
}

func makeValue(desiredVal int, nums []int) bool {
	operators := genOperators(len(nums) - 1)

	for _, ops := range operators {
		if evalP2(nums, ops) == desiredVal {
			return true
		}
	}
	return false
}

func solvePart2(arr []string) int {
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

		if makeValue(desiredVal, nums) {
			total += desiredVal
		}
	}

	return total
}
