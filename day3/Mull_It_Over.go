package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var (
	matchingRegx = `mul\((\d{1,3}),(\d{1,3})\)`
	removeMatch  = `(?s)don't\(\).*?(?:do\(\)|$)`
)

func main() {
	// go part1()
	part2()
}

func part2() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	inputData := string(content)

	removeRe := regexp.MustCompile(removeMatch)
	inputData = removeRe.ReplaceAllString(inputData, "")

	reMul := regexp.MustCompile(matchingRegx)
	matches := reMul.FindAllStringSubmatch(inputData, -1)

	sum := 0
	for _, match := range matches {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		sum += a * b
	}

	fmt.Println(sum)

}

// func part1() {
// 	context, err := os.ReadFile("input.txt")
// 	if err != nil {
// 		panic(err)
// 	}

// 	re := regexp.MustCompile(matchingRegx)

// 	matchs := re.FindAllStringSubmatch(string(context), -1)

// 	sum := 0
// 	for _, data := range matchs {
// 		a, _ := strconv.Atoi(data[1])
// 		b, _ := strconv.Atoi(data[2])
// 		sum += a * b
// 	}

// 	fmt.Println(sum)
// }
