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
	mult_r, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	num_r, _ := regexp.Compile(`\d{1,3}`)
	operations := mult_r.FindAllString(data, -1)
	for i := 0; i < len(operations); i++ {
		numbers := num_r.FindAllString(operations[i], 2)
		num_1, _ := strconv.Atoi(numbers[0])
		num_2, _ := strconv.Atoi(numbers[1])
		part_1 += num_1 * num_2
	}

	part_2_r, _ := regexp.Compile(`((mul\(\d{1,3},\d{1,3}\))|(do\(\))|(don't\(\)))`)
	operations = part_2_r.FindAllString(data, -1)
	do := true
	for i := 0; i < len(operations); i++ {
		if operations[i] == "do()" {
			do = true
		}
		if operations[i] == "don't()" {
			do = false
		}
		if do && mult_r.MatchString(operations[i]) {
			numbers := num_r.FindAllString(operations[i], 2)
			num_1, _ := strconv.Atoi(numbers[0])
			num_2, _ := strconv.Atoi(numbers[1])
			part_2 += num_1 * num_2
		}
	}

	fmt.Println(part_1, part_2)
}
