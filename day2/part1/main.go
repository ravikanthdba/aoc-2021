package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var depth, horizontalPosition int
	f, err := os.Open("input.txt")
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
		}

		if strings.ToLower(strings.Split(scanner.Text(), " ")[0]) == "up" {
			depth -= value
		}

		if strings.ToLower(strings.Split(scanner.Text(), " ")[0]) == "down" {
			depth += value
		}
	}

	fmt.Printf("The horizontal position is: %d\n", horizontalPosition*depth)
}
