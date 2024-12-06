package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	part1()
	part2()
}

func readFile() [][]rune {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, []rune(line))
	}
	return input
}

var direction = [8][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
	{-1, -1},
	{-1, 1},
	{1, -1},
	{1, 1},
}

func checkBounds(rows, cols, x, y int) bool {
	return x >= 0 && x < rows && y >= 0 && y < cols
}

func searchWord(grid [][]rune, word string) int {
	rows := len(grid)
	cols := len(grid[0])
	matches := 0

	for x, row := range grid {
		for y := range row {
			for _, d := range direction {
				match := true
				for i := 0; i < len(word); i++ {
					newX := x + i*d[0]
					newY := y + i*d[1]
					if !checkBounds(rows, cols, newX, newY) || grid[newX][newY] != rune(word[i]) {
						match = false
						break
					}
				}
				if match {
					matches++
				}
			}
		}
	}
	return matches
}

func part1() {
	grid := readFile()
	matches := searchWord(grid, "XMAS")
	fmt.Println("Part 1 Answer: ", matches)
}

func part2() {
	grid := readFile()
	matches := searchXmas(grid)
	fmt.Println("Part 2 Answer: ", matches)
}

func checkXmasPattern(topleft, topright, bottomleft, bottomright rune) bool {
	return (topleft == 'M' && topright == 'M' && bottomleft == 'S' && bottomright == 'S') ||
		(topleft == 'S' && topright == 'S' && bottomleft == 'M' && bottomright == 'M') ||
		(topleft == 'M' && topright == 'S' && bottomleft == 'M' && bottomright == 'S') ||
		(topleft == 'S' && topright == 'M' && bottomleft == 'S' && bottomright == 'M')
}

func searchXmas(grid [][]rune) int {
	rows := len(grid)
	cols := len(grid[0])
	matches := 0

	for x, row := range grid {
		for y := range row {
			if grid[x][y] != 'A' {
				continue
			}

			topleftX, topleftY := x-1, y-1
			toprightX, toprightY := x-1, y+1
			bottomleftX, bottomleftY := x+1, y-1
			bottomrightX, bottomrightY := x+1, y+1

			if !checkBounds(rows, cols, topleftX, topleftY) ||
				!checkBounds(rows, cols, toprightX, toprightY) ||
				!checkBounds(rows, cols, bottomleftX, bottomleftY) ||
				!checkBounds(rows, cols, bottomrightX, bottomrightY) {
				continue
			}

			topleft := grid[topleftX][topleftY]
			topright := grid[toprightX][toprightY]
			bottomleft := grid[bottomleftX][bottomleftY]
			bottomright := grid[bottomrightX][bottomrightY]

			if checkXmasPattern(topleft, topright, bottomleft, bottomright) {
				matches++
			}
		}
	}
	return matches
}
