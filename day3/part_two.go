package day3

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DayThreePartTwo() {
	file := "day3/input.txt"
	f, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("error reading file: %v\n", err)
		return
	}

	var total int64 = 0
	lines := strings.Split(string(f), "\n")
	cleanLines := lines[:len(lines)-1]

	for _, line := range cleanLines {
		vector := make([]byte, 0, 12)

		numbersLeft := len(line)
		for i := 0; i < len(line); i++ {
			actual := line[i]

			for len(vector) > 0 && vector[len(vector)-1] < actual && (len(vector)-1+numbersLeft) >= 12 {
				vector = vector[:len(vector)-1]
			}
			if len(vector) < 12 {
				vector = append(vector, actual)
			}
			numbersLeft--
		}
		result := string(vector)

		localSumInt, _ := strconv.Atoi(string(result))
		total += int64(localSumInt)
	}
	fmt.Println("total: ", total)
}
