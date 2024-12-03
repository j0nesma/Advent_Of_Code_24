package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"Advent_Of_Code_24/Error"
	"Advent_Of_Code_24/FileReader"
)

func main() {
	part1()
	part2()
}

func part1() {
	toString, err := FileReader.ReadFileToString("./Day3/input.txt")
	//toString, err := FileReader.ReadFileToString("./Day3/test.txt")
	Error.Check(err)
	var regex = `mul\([0-9][0-9]?[0-9]?\,[0-9][0-9]?[0-9]?\)`
	compile, err := regexp.Compile(regex)
	Error.Check(err)
	allString := compile.FindAllString(toString, -1)
	total := 0
	for i := range allString {
		split := strings.Split(allString[i], ",")
		a, err := strconv.Atoi(strings.Split(split[0], "mul(")[1])
		Error.Check(err)
		b, err := strconv.Atoi(strings.Split(split[1], ")")[0])
		total += a * b
	}
	fmt.Println(total)
}

func part2() {
	toString, err := FileReader.ReadFileToString("./Day3/input.txt")
	//toString, err := FileReader.ReadFileToString("./Day3/test2.txt")
	Error.Check(err)

	valid := ""
	DONT := "don't()"
	DO := "do()"
	for {
		if strings.Index(toString, DONT) == -1 {
			valid += toString
			break
		}
		valid += toString[:strings.Index(toString, DONT)]
		toString = toString[strings.Index(toString, DONT)+7:]
		if strings.Index(toString, DO) == -1 {
			break
		}
		toString = toString[strings.Index(toString, DO)+4:]
	}
	fmt.Println(valid)
	toString = valid
	regex := `mul\([0-9][0-9]?[0-9]?\,[0-9][0-9]?[0-9]?\)`
	compile, err := regexp.Compile(regex)
	Error.Check(err)
	allString := compile.FindAllString(toString, -1)
	total := 0
	for i := range allString {
		split := strings.Split(allString[i], ",")
		a, err := strconv.Atoi(strings.Split(split[0], "mul(")[1])
		Error.Check(err)
		b, err := strconv.Atoi(strings.Split(split[1], ")")[0])
		total += a * b
	}
	fmt.Println(total)
}
