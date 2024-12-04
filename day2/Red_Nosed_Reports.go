package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() int {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safe := 0
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)

		intNums, err := convertToIntNums(nums)
		if err != nil {
			panic(err)
		}

		increasing := intNums[0] <= intNums[1]
		changed := false
		for i := 1; i < len(intNums); i++ {
			diff := intNums[i] - intNums[i-1]
			if (diff < 0 && increasing) || (diff > 0 && !increasing) || abs(diff) > 3 || diff == 0 {
				changed = true
				break
			}

		}
		if !changed {
			safe++
		}
	}
	return safe
}

func part2() int {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safe := 0
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)

		reports, err := convertToIntNums(nums)
		if err != nil {
			panic(err)
		}

		if isSafe(reports) {
			safe++
		} else {
			for level := range reports {
				temp := append(append([]int{}, reports[:level]...), reports[level+1:]...)
				if isSafe(temp) {
					safe++
					break
				}
			}
		}

	}
	return safe
}

func isSafe(r []int) bool {
	if r[0] > r[1] {
		for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
	}
	for i := 0; i < len(r)-1; i++ {
		if r[i+1]-r[i] < 1 || r[i+1]-r[i] > 3 {
			return false
		}
	}
	return true
}

func convertToIntNums(nums []string) ([]int, error) {
	intNums := make([]int, len(nums))
	for i, num := range nums {
		var err error
		intNums[i], err = strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
	}
	return intNums, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
