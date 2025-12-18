package day6

import (
	"bytes"
	"fmt"
	"os"
)

func DaySixPartTwo() {
	file, err := os.ReadFile("day6/input.txt")
	if err != nil {
		fmt.Printf("error reading file: %v\n", err)
		return
	}

	lines := bytes.Split(file, []byte("\n"))
	data := make([][]byte, 0)
	for _, line := range lines {
		if len(line) > 0 {
			data = append(data, line)
		}
	}

	var res int64
	var currP []int64
	var op func(a int64, b int64) int64

	x := 0
outer:
	for {
		curr := int64(0)
		allSpaces := true
		for y := 0; y < len(data); y++ {
			if x >= len(data[y]) {
				break outer
			}

			switch data[y][x] {
			case ' ', '\n':
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				allSpaces = false
				curr = curr*10 + int64(data[y][x]-'0')
			case '+':
				allSpaces = false
				op = func(a int64, b int64) int64 {
					return a + b
				}
			case '*':
				allSpaces = false
				op = func(a int64, b int64) int64 {
					return a * b
				}
			}
		}
		if allSpaces {
			r := currP[0]
			for i := 1; i < len(currP); i++ {
				r = op(r, currP[i])
			}
			res += r
			currP = currP[:0]
		} else {
			currP = append(currP, curr)
		}
		x++
	}

	r := currP[0]
	for i := 1; i < len(currP); i++ {
		r = op(r, currP[i])
	}
	res += r

	fmt.Println(res)
}
