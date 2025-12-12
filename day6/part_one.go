package day6

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DaySixPartOne() {
	f, err := os.ReadFile("./day6/input.txt")
	if err != nil {
		fmt.Printf("error reading file: %v\n", err)
		return
	}

	lines := strings.Split(string(f), "\n")
	var elems [][]string

	for _, line := range lines {
		row := strings.Fields(line)
		if len(row) == 0 {
			continue
		}
		elems = append(elems, row)
	}

	operations := elems[len(elems)-1]
	numbers := elems[:len(elems)-1]
	nCols := len(numbers[0])

	var res int64 = 0
	for i := range nCols {
		var localSum int64 = 0
		oper := operations[i]

		// otherwise always 0
		if oper == "*" {
			localSum = 1
		}

		for _, row := range numbers {
			vInt, _ := strconv.Atoi(row[i])
			if oper == "+" {
				localSum += int64(vInt)
			} else {
				localSum *= int64(vInt)
			}
		}
		res += localSum
	}

	fmt.Println("sum: ", res)
}
