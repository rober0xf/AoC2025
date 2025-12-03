package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DayOnePartTwo() {
	file := "day1/input.txt"
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("error reading the file: %v", err)
		return
	}
	parts := strings.Split(strings.TrimSpace(string(content)), "\n")

	total := 0
	index := 50

	for _, v := range parts {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}

		side := v[0]
		number, _ := strconv.Atoi(v[1:])

		for i := 0; i < number; i++ {
			if side == 'L' {
				index--
				if index < 0 {
					index = 99
				}
			} else {
				index++
				if index >= 100 {
					index = 0
				}
			}
			if index == 0 {
				total++
			}
		}
	}
	fmt.Println("total: ", total)
}
