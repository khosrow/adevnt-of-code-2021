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

	return diag, nil
}

func analyzeDiagnostics(input []string) (string, error) {
	aggregate := ""

	for i := 0; i < len(input[0]); i++ {
		num := 0
		for _, sig := range input {
			if i >= len(sig) {
				continue
			}
			n, err := strconv.Atoi(string(sig[i]))
			if err != nil {
				return aggregate, err
			}
			num += n
		}
		if num >= len(input)/2 {
			aggregate += "1"
		} else {
			aggregate += "0"
		}
	}

	return aggregate, nil
}

func reverseAggregate(aggregate string, mask uint64, bits int) (string, error) {
	aggUint, err := strconv.ParseUint(aggregate, 2, bits)
	if err != nil {
		return "", err
	}

	revAggString := strconv.FormatUint(mask^aggUint, 2)
	if len(revAggString) != bits {
		for i := 0; i < bits-len(revAggString); i++ {
			revAggString = "0" + revAggString
		}
	}

	fmt.Printf("Agg:\t %s\n", aggregate)
	fmt.Printf("Rev Agg: %s\n", revAggString)

	return revAggString, nil
}

func getPowerConsumption(aggregate string, mask uint64) (uint64, error) {
	gamma, err := strconv.ParseUint(aggregate, 2, 64)
	if err != nil {
		return 0, err
	}
	epsilon := mask ^ gamma

	fmt.Printf("gamma: %d, %s\n", gamma, strconv.FormatUint(uint64(gamma), 2))
	fmt.Printf("epsilon: %d, %s\n", epsilon, strconv.FormatUint(uint64(epsilon), 2))

	return gamma * epsilon, nil

}

func getLifeSupportRating(aggregate string, input []string) (uint64, error) {
	var tempBuffer []string

	buffer := input
	for i := 0; i < len(aggregate); i++ {
		if len(buffer) == 1 {
			break
		}
		for _, sig := range buffer {
			if i >= len(sig) {
				continue
			}
			if string(sig[i]) == string(aggregate[i]) {
				fmt.Printf("aggregate: %s, signal: %s, bit number: %d\n", aggregate, sig, i)
				tempBuffer = append(tempBuffer, sig)
			}
		}
		if len(tempBuffer) != 0 {
			buffer = tempBuffer
		}
		tempBuffer = nil
	}

	if len(buffer) != 1 {
		return 0, fmt.Errorf("something has gone wrong. Buffer length is 0")
	}
	rating, err := strconv.ParseUint(buffer[0], 2, 64)
	if err != nil {
		return 0, err
	}
	fmt.Printf("Rating: %d, %s\n", rating, buffer[0])

	return rating, nil
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

	mask, err := strconv.ParseUint(strings.Repeat("1", len(diag[0])), 2, len(diag[0]))
	aggregate, err := analyzeDiagnostics(diag)
	if err != nil {
		panic(err)
	}

	reverse_aggregate, err := reverseAggregate(aggregate, mask, len(diag[0]))
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}
	powerConsumption, err := getPowerConsumption(aggregate, mask)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Power Consumption: %d\n", powerConsumption)

	oxygen, err := getLifeSupportRating(aggregate, diag)
	if err != nil {
		panic(err)
	}

	co2, err := getLifeSupportRating(reverse_aggregate, diag)
	if err != nil {
		panic(err)
	}

	fmt.Printf("oxygen: %d, co2: %d\n", oxygen, co2)
	fmt.Printf("Life support rating: %d\n", oxygen*co2)
}
