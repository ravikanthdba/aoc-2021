package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var aim, horizontalPosition, depth int
	f, err := os.Open("/Users/rgarimel/Documents/Programming/aoc-2021/day2/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value, err := strconv.Atoi(strings.Split(scanner.Text(), " ")[1])
		if err != nil {
			fmt.Println(err)
			return
		}

		if strings.ToLower(strings.Split(scanner.Text(), " ")[0]) == "forward" {
			horizontalPosition += value
			depth = depth + (aim * value)
		}

		if strings.ToLower(strings.Split(scanner.Text(), " ")[0]) == "up" {
			aim -= value
		}

		if strings.ToLower(strings.Split(scanner.Text(), " ")[0]) == "down" {
			aim += value
		}

	}

	fmt.Printf("The actual position of the submarine is: %d\n", horizontalPosition*depth)
}
