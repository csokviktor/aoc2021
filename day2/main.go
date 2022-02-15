package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type commands struct {
	direction string
	amount    int
}

func readInput(filepath string) []commands {
	bytes, _ := ioutil.ReadFile(filepath)
	stringData := string(bytes)
	stringLines := strings.Split(stringData, "\n")
	vals := make([]commands, len(stringLines))
	for i, value := range stringLines {
		splits := strings.Split(value, " ")
		vals[i].direction = splits[0]
		//fmt.Println(strconv.Atoi(strings.TrimSuffix(splits[1], "\r")))
		vals[i].amount, _ = strconv.Atoi(strings.TrimSuffix(splits[1], "\r"))
	}
	return vals
}

func solve1(input []commands) int {
	hor := 0
	ver := 0
	for _, element := range input {
		switch element.direction {
		case "forward":
			hor += element.amount
		case "up":
			ver -= element.amount
		case "down":
			ver += element.amount
		}
	}
	return hor * ver
}

func solve2(input []commands) int {
	hor := 0
	ver := 0
	aim := 0
	for _, element := range input {
		switch element.direction {
		case "forward":
			hor += element.amount
			ver += aim * element.amount
		case "up":
			aim -= element.amount
		case "down":
			aim += element.amount
		}
	}
	return hor * ver
}

func main() {
	input := readInput("input.txt")
	res1 := solve1(input)
	fmt.Println(res1)
	res2 := solve2(input)
	fmt.Println(res2)
}
