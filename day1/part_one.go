package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DayOnePartOne() {
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
		switch v[0] {
		case 'L':
			lnumber, _ := strconv.Atoi(v[1:])
			pos := ((index-lnumber)%100 + 100) % 100 // fix negative numbers
			if pos == 0 {
				total++
			}
			index = pos // reset the position
		case 'R':
			rnumber, _ := strconv.Atoi(v[1:])
			pos := (index + rnumber) % 100
			if pos == 0 {
				total++
			}
			index = pos
		}
	}
	fmt.Println("total: ", total)

}
