package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	part_1, part_2 := 0, 0
	raw_data, _ := os.ReadFile("input.txt")
	data := string(raw_data)

	// Regex to get mult(x,y), do(), and don't()
	operations_r, _ := regexp.Compile(`((mul\(\d{1,3},\d{1,3}\))|(do\(\))|(don't\(\)))`)

	// Regex just for mult(x,y)
	mult_r, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)

	// Regex for digits in mult(x,y)
	num_r, _ := regexp.Compile(`\d{1,3}`)

	// Get all the possible ops
	operations := operations_r.FindAllString(data, -1)

	// Switch for part 2
	do := true
	for i := 0; i < len(operations); i++ {
		if operations[i] == "do()" {
			do = true
		}
		if operations[i] == "don't()" {
			do = false
		}
		// If it's a mult(x,y) then work on it
		if mult_r.MatchString(operations[i]) {
			numbers := num_r.FindAllString(operations[i], 2)
			num_1, _ := strconv.Atoi(numbers[0])
			num_2, _ := strconv.Atoi(numbers[1])
			prod := num_1 * num_2
			part_1 += prod
			if do {
				part_2 += prod
			}
		}
	}

	fmt.Println(part_1, part_2)
}
