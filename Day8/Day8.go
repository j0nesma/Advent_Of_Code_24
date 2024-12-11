package main

import (
	"fmt"
	"slices"
	"strings"

	"Advent_Of_Code_24/Error"
	"Advent_Of_Code_24/FileReader"
)

type pos struct {
	X, Y float64
}

func main() {
	//toString, err := FileReader.ReadFileToString("./Day8/test.txt")
	toString, err := FileReader.ReadFileToString("./Day8/input.txt")
	Error.Check(err)

	rows := strings.Split(toString, "\n")
	m := make(map[string][]pos)
	for i, row := range rows {
		row = strings.TrimSpace(row)
		values := strings.Split(row, "")
		for j, value := range values {
			if value == "." {
				continue
			}
			m[value] = append(m[value], pos{float64(i), float64(j)})
		}
	}

	count := 0
	signals := make(map[string][]pos)
	for k, v := range m {
		for i := range v {
			for j := range v {
				if i == j {
					continue
				}
				val1 := v[i]
				val2 := v[j]

				diffY := val1.Y - val2.Y
				diffX := val1.X - val2.X

				if diffY == 0 && diffX == 0 {
					continue
				}
				position := pos{X: val1.X + diffX, Y: val1.Y + diffY}
				for {
					signals[k] = append(signals[k], position)
					if position.X < 0 || position.Y < 0 || position.X > float64(len(rows)-1) || position.Y > float64(len(rows[0])-2) {
						break
					}
					position = pos{X: position.X + diffX, Y: position.Y + diffY}
				}
				position = pos{X: val1.X - diffX, Y: val1.Y - diffY}
				for {
					signals[k] = append(signals[k], position)
					if position.X < 0 || position.Y < 0 || position.X > float64(len(rows)-1) || position.Y > float64(len(rows[0])-2) {
						break
					}
					position = pos{X: position.X - diffX, Y: position.Y - diffY}
				}
			}
		}
	}
	vPoints := []pos{}
	for _, s := range signals {
		for _, v := range s {
			// Y is weird and has to have 2 removed from it to make it pass
			if v.X < 0 || v.Y < 0 || v.X > float64(len(rows)-1) || v.Y > float64(len(rows[0])-2) {
				continue
			}
			if slices.Contains(vPoints, v) {
				continue
			}
			vPoints = append(vPoints, v)
			count++
		}
	}
	fmt.Printf("rows length = %v, columns length = %v\n", len(rows), len(rows[0]))
	fmt.Printf("%v\n", vPoints)
	fmt.Println(count)
}
