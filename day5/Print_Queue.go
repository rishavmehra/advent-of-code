package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	Scanner := bufio.NewScanner(file)
	edge := map[int][]int{}
	for Scanner.Scan() {
		text := Scanner.Text()
		if text == "" {
			break
		}
		src, dist := getEdge(text)
		edge[src] = append(edge[src], dist)
	}

	updates := [][]int{}
	for Scanner.Scan() {
		text := Scanner.Text()
		updates = append(updates, getUpdate(text))
	}

	total := 0
	for _, update := range updates {
		path := map[int]int{}
		fixed := false

		for i := 0; i <= len(update)-1; i++ {
			page := update[i]
			path[page] = i
			for _, neighbor := range edge[page] {
				if j, ok := path[neighbor]; ok && path[neighbor] < path[page] {
					fixed = true
					update[i], update[j] = update[j], update[i]
					path[neighbor], path[page] = path[page], path[neighbor]
					i = path[page]
				}
			}
		}
		if fixed {
			total += update[len(update)/2]
		}
	}
	fmt.Println("Part2: ", total)
}

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	Scanner := bufio.NewScanner(file)
	edge := map[int][]int{}
	for Scanner.Scan() {
		text := Scanner.Text()
		if text == "" {
			break
		}
		src, dist := getEdge(text)
		edge[src] = append(edge[src], dist)
	}

	updates := [][]int{}
	for Scanner.Scan() {
		text := Scanner.Text()
		updates = append(updates, getUpdate(text))
	}

	total := 0

	for _, update := range updates {
		path := map[int]bool{}

	CHECK:
		for _, page := range update {
			for _, neighbor := range edge[page] {
				if _, ok := path[neighbor]; ok {
					break CHECK
				}
			}
			path[page] = true
		}
		if len(path) == len(update) {
			total += update[len(update)/2]
		}
	}
	fmt.Println("Part1: ", total)
}

func getEdge(text string) (int, int) {
	parts := strings.Split(text, "|")
	src, _ := strconv.Atoi(parts[0])
	dist, _ := strconv.Atoi(parts[1])
	return src, dist
}

func getUpdate(text string) []int {
	res := []int{}
	parts := strings.Split(text, ",")
	for _, part := range parts {
		val, _ := strconv.Atoi(part)
		res = append(res, val)
	}
	return res
}
