package day3

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DayThreePartOne() {
	file := "day3/input.txt"
	f, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("error reading file: %v\n", err)
		return
	}

	total := 0
	lines := strings.Split(string(f), "\n")
	cleanLines := lines[:len(lines)-1]

	for _, line := range cleanLines {
		var localSum string
		largest := 0
		secondLargest := 0

		var idx int
		for i, v := range line {
			if i == len(line)-1 {
				break
			} else {
				if int(v) > largest {
					largest = int(v)
					idx = i
				}
			}
		}

		for i, v := range line {
			if i != idx {
				if int(v) > secondLargest && i > idx {
					secondLargest = int(v)
				}
			}
		}
		localSum = string(largest) + string(secondLargest)
		fmt.Println(localSum)
		localSumInt, _ := strconv.Atoi(string(localSum))
		total += localSumInt
	}
	fmt.Println("total: ", total)
}
