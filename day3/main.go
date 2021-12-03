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

func main() {
	var gammaRateBits, epsilonRateBits string
	f, err := os.Open("input.txt")
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

	// var bitsList []string
	for i := 0; i < recordLength; i++ {
		var bits string
		for _, rec := range input {
			bits += string(rec[i])
		}
		gammaRateBits += checkMaxOccurredBit(bits)
		epsilonRateBits += checkLeastOccurredBit(bits)
	}

	gammaNumber, _ := convertBinaryToNumber(gammaRateBits)
	epsilonNumber, _ := convertBinaryToNumber(epsilonRateBits)
	fmt.Println(gammaNumber * epsilonNumber)
}
