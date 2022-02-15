package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func readInput(fileName string) []string {
	bytes, _ := ioutil.ReadFile(fileName)
	stringData := string(bytes)
	stringLines := strings.Split(string(stringData), "\n")
	return stringLines
}

func solve1(input []string) int {
	sums := make([]int, len(input[0])-1)
	for _, line := range input {
		for i, character := range line {
			if string(character) == "1" {
				sums[i] += 1
			}
		}
	}

	g := 0
	e := 0
	for j, sum := range sums {
		pos := len(sums) - j - 1
		if sum > len(input)/2 {
			g += int(math.Pow(float64(2), float64(pos)))
		} else {
			e += int(math.Pow(float64(2), float64(pos)))
		}
	}
	return e * g
}

func createSum(input []string) []int {
	sums1 := make([]int, len(input[0])-1)
	for _, line := range input {
		for i, character := range line {
			if string(character) == "1" {
				sums1[i] += 1
			}
		}
	}
	return sums1
}

func solve2(input []string) int64 {
	currentLines1 := input
	currentLines2 := input

	for i := 0; i < len(input[0])-1; i++ {
		if len(currentLines1) > 1 {
			sums1 := createSum(currentLines1)
			var newCurr1 []string
			for _, element := range currentLines1 {
				if float64(sums1[i]) >= float64(len(currentLines1))/2.0 && string(element[i]) == "1" {
					newCurr1 = append(newCurr1, element)
				} else if float64(sums1[i]) < float64(len(currentLines1))/2.0 && string(element[i]) == "0" {
					newCurr1 = append(newCurr1, element)
				}
			}
			currentLines1 = newCurr1
		}

		if len(currentLines2) > 1 {
			sums2 := createSum(currentLines2)
			var newCurr2 []string
			for _, element := range currentLines2 {
				if float64(len(currentLines2)-sums2[i]) <= float64(len(currentLines2))/2.0 && string(element[i]) == "0" {
					newCurr2 = append(newCurr2, element)
				} else if float64(len(currentLines2)-sums2[i]) > float64(len(currentLines2))/2.0 && string(element[i]) == "1" {
					newCurr2 = append(newCurr2, element)
				}
			}
			currentLines2 = newCurr2
		}

	}
	ox, _ := strconv.ParseInt(strings.TrimSuffix(currentLines1[0], "\r"), 2, 64)
	src, _ := strconv.ParseInt(strings.TrimSuffix(currentLines2[0], "\r"), 2, 64)
	return ox * src

}

func main() {
	input := readInput("input.txt")
	res1 := solve1(input)
	fmt.Println(res1)
	res2 := solve2(input)
	fmt.Println(res2)
}
