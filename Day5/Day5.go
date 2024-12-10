package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"Advent_Of_Code_24/FileReader"
)

type Rule struct {
	before, after int
}

func parseRules(lines []string) ([]Rule, int) {
	var rules []Rule
	emptyLineIdx := 0

	for i, line := range lines {
		if line == "" {
			emptyLineIdx = i
			break
		}
		parts := strings.Split(strings.TrimSpace(line), "|")
		before, _ := strconv.Atoi(parts[0])
		after, _ := strconv.Atoi(parts[1])
		rules = append(rules, Rule{before, after})
	}
	return rules, emptyLineIdx
}

func parseUpdate(line string) []int {
	parts := strings.Split(strings.TrimSpace(line), ",")
	update := make([]int, len(parts))
	for i, p := range parts {
		update[i], _ = strconv.Atoi(strings.TrimSpace(p))
	}
	return update
}

func isValidOrder(update []int, rules []Rule) bool {
	positions := make(map[int]int)
	for i, page := range update {
		positions[page] = i
	}

	for _, rule := range rules {
		beforePos, beforeExists := positions[rule.before]
		afterPos, afterExists := positions[rule.after]

		if beforeExists && afterExists && beforePos > afterPos {
			return false
		}
	}
	return true
}

func solvePart1(input []string) int {
	rules, emptyLineIdx := parseRules(input)
	sum := 0

	for _, line := range input[emptyLineIdx+1:] {
		if line == "" {
			continue
		}
		update := parseUpdate(line)
		if isValidOrder(update, rules) {
			middleIdx := len(update) / 2
			sum += update[middleIdx]
		}
	}
	return sum
}
func sortPages(pages []int, rules []Rule) []int {
	sorted := make([]int, len(pages))
	copy(sorted, pages)

	sort.Slice(sorted, func(i, j int) bool {
		a, b := sorted[i], sorted[j]
		// Check if there's a rule requiring b before a
		for _, rule := range rules {
			if rule.before == b && rule.after == a {
				return false
			}
			if rule.before == a && rule.after == b {
				return true
			}
		}
		// If no direct rule, maintain current order
		return i < j
	})

	return sorted
}

func solvePart2(input []string) int {
	rules, emptyLineIdx := parseRules(input)
	sum := 0

	for _, line := range input[emptyLineIdx+1:] {
		if line == "" {
			continue
		}
		update := parseUpdate(line)
		if !isValidOrder(update, rules) {
			sortedUpdate := sortPages(update, rules)
			middleIdx := len(sortedUpdate) / 2
			sum += sortedUpdate[middleIdx]
		}
	}
	return sum
}

func main() {
	//val, _ := FileReader.ReadFileToString("./Day5/test.txt")
	val, _ := FileReader.ReadFileToString("./Day5/input.txt")
	input := strings.Split(val, "\r\n")

	fmt.Println("Part 1 Result:", solvePart1(input))
	fmt.Println("Part 2 Result:", solvePart2(input))
}
