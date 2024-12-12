package main

import (
	"fmt"
	"strconv"
	"strings"

	"Advent_Of_Code_24/Error"
	"Advent_Of_Code_24/FileReader"
)

func main() {
	toString, err := FileReader.ReadFileToString("./Day9/test.txt")
	//toString, err := FileReader.ReadFileToString("./Day9/input.txt")
	Error.Check(err)
	digits := strings.Split(toString, "")
	fullString := ""
	for i, digit := range digits {
		atoi, _ := strconv.Atoi(digit)
		for j := 0; j < atoi; j++ {
			if i%2 == 0 {
				fullString += strconv.Itoa(i / 2)
			} else {
				fullString += "."
			}
			if i < len(digits) {
				fullString += ","
			}
		}
	}
	fullString = fullString[:len(fullString)-1]
	fmt.Println(fullString)
	disk := part1(fullString)
	fmt.Println(disk)
	disk2 := part2(fullString)
	fmt.Println(disk2)
	count := 0
	for i := 0; i < len(disk); i++ {
		if disk[i] == "." {
			break
		}
		atoi, err := strconv.Atoi(disk[i])
		Error.Check(err)
		fmt.Printf("%v * %v = %v\n", atoi, i, i*atoi)
		count += atoi * i
	}
	fmt.Println()
	fmt.Println(count)
}

// TODO
func part2(fullString string) []string {
	disk := strings.Split(fullString, ",")
	complete := false
	for i := len(disk) - 1; i > 0; i-- {
		if disk[i] == "." {
			continue
		}
		numberOf := 1
		k := i - 1
		for {
			if k == -1 {
				break
			}
			if disk[k] == disk[i] {
				numberOf++
			} else {
				break
			}
			k--
		}
		println(disk[i], numberOf)
		for j := 0; j < len(disk); j++ {
			if j >= i {
				complete = true
				break
			}
			if disk[j] != "." {
				continue
			}
			countOf := 1
			k := j + 1
			for {
				if disk[k] == disk[j] {
					numberOf++
				} else {
					break
				}
				k++
			}
			if countOf < numberOf {
				break
			}
			for l := j; l < j+numberOf; l++ {
				temp := disk[l]
				disk[l] = disk[i]
				disk[i] = temp
			}
			break
		}
		if complete {
			break
		}
		i -= i - numberOf
	}
	return disk
}

func part1(fullString string) []string {
	disk := strings.Split(fullString, ",")
	for i := len(disk) - 1; i > 0; i-- {
		complete := false
		if disk[i] == "." {
			continue
		}
		for j := 0; j < len(disk); j++ {
			if j >= i {
				complete = true
				break
			}
			if disk[j] != "." {
				continue
			}

			temp := disk[j]
			disk[j] = disk[i]
			disk[i] = temp
			break
		}
		if complete {
			break
		}
	}
	return disk
}
