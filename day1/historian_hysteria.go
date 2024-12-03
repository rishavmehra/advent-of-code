package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {

	start := time.Now()
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("not able to open the file %s: %v", file.Name(), err)
	}

	var left []int
	var right []int
	var temp int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		temp, err = strconv.Atoi(line[0:5])
		left = append(left, temp)
		temp, err = strconv.Atoi(line[8:13])
		right = append(right, temp)
	}
	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})
	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	var dist_temp int = 0
	var dist int
	var right_index int = 0
	var sim_score int = 0
	var occs int
	for i := 0; i < len(left); i++ {
		dist_temp = left[i] - right[i]
		if dist_temp < 0 {
			dist_temp = -dist_temp
		}
		dist += dist_temp

		occs = 0
		var dist_temp2 int = left[i] - right[right_index]
		if dist_temp2 >= 0 {
			for dist_temp2 >= 0 {
				if dist_temp2 == 0 {
					occs += 1
				}
				if right_index < len(left)-1 {
					right_index += 1
					dist_temp2 = left[i] - right[right_index]
				} else {
					break
				}
			}
			sim_score += left[i] * occs
		}

	}

	fmt.Printf("%d, %d\n", dist, sim_score)
	elapsed := time.Since(start)
	fmt.Printf("%v\n", elapsed)

}
