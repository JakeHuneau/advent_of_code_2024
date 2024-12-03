package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part_1, part_2 := 0, 0

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		raw_numbers := strings.Split(scanner.Text(), " ")
		var numbers []int
		for i := 0; i < len(raw_numbers); i++ {
			num, _ := strconv.Atoi(raw_numbers[i])
			numbers = append(numbers, num)
		}
		if is_test_safe(numbers) {
			part_1++
		}
		fmt.Println("Testing ", numbers)
		for i := 0; i < len(numbers); i++ {
			var sliced_numbers []int
			sliced_numbers = append(sliced_numbers, numbers[:i]...)
			sliced_numbers = append(sliced_numbers, numbers[i+1:]...)
			fmt.Println(sliced_numbers)
			if is_test_safe(sliced_numbers) {
				fmt.Println("Safe")
				part_2++
				break
			}
		}
	}

	fmt.Println(part_1, part_2)
}

func is_test_safe(line []int) bool {
	safe_difference := 3
	is_increasing := line[1] > line[0]
	for i := 0; i < len(line)-1; i++ {
		num_1 := line[i]
		num_2 := line[i+1]
		if num_1 == num_2 {
			return false
		}

		if num_2 < num_1 && is_increasing {
			return false
		}
		if num_2 > num_1 && !is_increasing {
			return false

		}
		if abs_diff(num_1, num_2) > safe_difference {
			return false

		}
	}
	return true
}

func abs_diff(num1 int, num2 int) int {
	diff := num1 - num2
	if diff < 0 {
		return -diff
	}
	return diff
}
