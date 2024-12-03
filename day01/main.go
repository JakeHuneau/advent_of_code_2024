package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var nums [2][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ")
		for i := 0; i < 2; i++ {
			num, err := strconv.Atoi(parts[i])
			if err != nil {
				panic(err)
			}
			nums[i] = append(nums[i], num)
		}
	}

	sort.Ints(nums[0])
	sort.Ints(nums[1])

	total_distance := 0
	for i := 0; i < len(nums[0]); i++ {
		total_distance += int(math.Abs(float64(nums[0][i] - nums[1][i])))
	}
	fmt.Println("Part 1:", total_distance)

	part_2_sol := 0
	for i := 0; i < len(nums[0]); i++ {
		test_value := nums[0][i]
		for j := 0; j < len(nums[1]); j++ {
			if nums[1][j] == test_value {
				part_2_sol += test_value
			}
		}
	}
	fmt.Println("Part 2:", part_2_sol)
}
