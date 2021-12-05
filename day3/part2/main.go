package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

//get length of all records are the same
func getLength(input []int) (int, error) {
	if len(input) == 0 {
		return 0, fmt.Errorf("%q", "no valid input")
	}

	comparisionValue := input[0]
	for i := 1; i < len(input); i++ {
		if input[i] != comparisionValue {
			return 0, nil
		}
	}
	return comparisionValue, nil
}

func checkMaxOccurredBit(bits string) string {
	var bitOccurance = make(map[string]int)
	for _, bit := range bits {
		bitOccurance[string(bit)]++
	}

	if bitOccurance["0"] > bitOccurance["1"] {
		return "0"
	}

	if bitOccurance["0"] == bitOccurance["1"] {
		return "1"
	}

	return "1"
}

func checkLeastOccurredBit(bits string) string {
	var bitOccurance = make(map[string]int)
	for _, bit := range bits {
		bitOccurance[string(bit)]++
	}

	if bitOccurance["0"] < bitOccurance["1"] {
		return "0"
	}

	if bitOccurance["0"] == bitOccurance["1"] {
		return "0"
	}

	return "1"
}

func convertBinaryToNumber(binary string) (int, error) {
	var number int
	binaryNumber, err := strconv.Atoi(binary)
	if err != nil {
		return 0, fmt.Errorf("%q", err)
	}
	var count int

	for binaryNumber != 0 {
		number = number + ((binaryNumber % 10) * int(math.Pow(2, float64(count))))
		count++
		binaryNumber = binaryNumber / 10
	}

	return number, nil
}

func getRating(input []string, recordLength int, rating string) (int, error) {
	for i := 0; i < recordLength; i++ {
		var bits string
		var output []string
		for _, record := range input {
			bits += string(record[i])
		}

		if rating != "oxygen" && rating != "co2" {
			return 0, fmt.Errorf("%q", "invalid rating specifier. It needs to be \"oxygen (or) \"co2")
		}

		switch rating {
		case "oxygen":
			checkMaxRecurredByte := checkMaxOccurredBit(bits)
			for _, record := range input {
				if string(record[i]) == checkMaxRecurredByte {
					output = append(output, record)
				}
			}

		case "co2":
			checkLeastRecurredByte := checkLeastOccurredBit(bits)
			for _, record := range input {
				if string(record[i]) == checkLeastRecurredByte {
					output = append(output, record)
				}
			}
		}

		input = output

		if len(input) == 1 {
			break
		}
	}

	oxygenGeneratorRating, err := convertBinaryToNumber(input[0])
	if err != nil {
		return 0, err
	}

	return oxygenGeneratorRating, nil
}

func main() {
	f, err := os.Open("/Users/rgarimel/Documents/Programming/aoc-2021/day3/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	var input []string

	scanner := bufio.NewScanner(f)
	var records []int
	for scanner.Scan() {
		input = append(input, scanner.Text())
		records = append(records, len(scanner.Text()))
	}

	recordLength, err := getLength(records)
	if err != nil {
		fmt.Println(err)
		return
	}

	oxygenGeneratorRating, err := getRating(input, recordLength, "oxygen")
	if err != nil {
		fmt.Println(err)
		return
	}

	co2ScrubberRating, err := getRating(input, recordLength, "co2")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Life Support Rating: ", oxygenGeneratorRating*co2ScrubberRating)
}
