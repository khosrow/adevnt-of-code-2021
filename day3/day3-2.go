package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type move struct {
	direction string
	amount    int
}

func parseDiagnostics(input string) ([]string, error) {
	var diag []string
	for _, line := range strings.Split(string(input), "\n") {
		diag = append(diag, line)
	}

	return diag[:len(diag)-1], nil
}

func analyzeBits(input []string, i int) (zeros int, ones int) {
	zeros = 0
	ones = 0
	for _, sig := range input {
		if string(sig[i]) == "0" {
			zeros += 1
		} else {
			ones += 1
		}
	}
	return zeros, ones
}

func checkBitCriteria(input []string, i int, criteria string) []string {
	var res []string
	for _, line := range input {
		if criteria == string(line[i]) {
			res = append(res, line)
		}
	}
	return res
}

func getOxygen(input []string) string {
	var tempBuffer []string
	buffer := input
	aggregate := ""

	for i := 0; i < len(input[0]); i++ {
		zeros, ones := analyzeBits(buffer, i)
		if ones >= zeros {
			aggregate = "1"
		} else {
			aggregate = "0"
		}
		tempBuffer = checkBitCriteria(buffer, i, aggregate)
		buffer = tempBuffer
		if len(tempBuffer) == 1 {
			break
		}
		tempBuffer = nil
	}
	return tempBuffer[0]
}

func getCO2(input []string) string {
	var tempBuffer []string

	buffer := input
	aggregate := ""

	for i := 0; i < len(input[0]); i++ {
		zeros, ones := analyzeBits(buffer, i)
		if zeros <= ones {
			aggregate = "0"
		} else {
			aggregate = "1"
		}

		tempBuffer = checkBitCriteria(buffer, i, aggregate)
		buffer = tempBuffer
		if len(tempBuffer) == 1 {
			break
		}
		tempBuffer = nil
	}
	return tempBuffer[0]
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	diag, err := parseDiagnostics(string(input))
	if err != nil {
		panic(err)
	}

	oxygenSt := getOxygen(diag)
	fmt.Printf("oxygen: %s\n", oxygenSt)

	co2St := getCO2(diag)
	fmt.Printf("CO2: %s\n", co2St)

	oxygen, _ := strconv.ParseUint(oxygenSt, 2, 64)
	co2, _ := strconv.ParseUint(co2St, 2, 64)
	fmt.Printf("Life support rating: %d\n", oxygen*co2)
}
