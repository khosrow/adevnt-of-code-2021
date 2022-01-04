package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Digit map[string]bool

type DisplayData struct {
	signals []string
	outputs []string
}

const (
	one   = "cf"
	four  = "bcdf"
	seven = "acf"
	eight = "abcdefg"
)

func buildDigit(input string) Digit {
	output := make(Digit)
	for _, s := range input {
		output[string(s)] = true
	}
	return output
}
func parseInput(input string) []DisplayData {
	var data []DisplayData
	for _, line := range strings.Split(input, "\n") {
		signals := make([]string, 0)
		outputs := make([]string, 0)
		if line == "" {
			continue
		}
		parts := strings.Split(line, " | ")
		for _, sig := range strings.Split(parts[0], " ") {
			signals = append(signals, sig)
		}

		for _, sig := range strings.Split(parts[1], " ") {
			outputs = append(outputs, sig)
		}
		data = append(data, DisplayData{
			signals: signals,
			outputs: outputs,
		})
	}
	return data
}

// subtract s2 from s1, then return result
// i.e s1 = "abcd", s2 = "bd"
// s1 - s2 = "ac"
func subtract(s1 string, s2 string) string {
	s := s1
	for _, c := range s2 {
		s = strings.ReplaceAll(s, string(c), "")
	}
	return s
}

// return combined elements of d1 and d2
func union(d1 Digit, d2 Digit) Digit {
	d3 := make(Digit)
	for k, v := range d1 {
		d3[k] = v
	}
	for k, v := range d2 {
		d3[k] = v
	}
	return d3
}

// func intersection(d1 Digit, d2 Digit) Digit {

// }

// Return all elements in s1 but not in s2
// i.e s1 = "abcd", s2 = "bd"
// s1 - s2 = "ac"
func difference(d1 Digit, d2 Digit) Digit {
	d3 := make(Digit)
	for k, v := range d1 {
		d3[k] = v
	}
	for k := range d3 {
		if _, ok := d2[k]; ok {
			delete(d3, k)
		}
	}
	return d3
}

func d2s(d Digit) string {
	s := ""
	for k := range d {
		s += string(k)
	}
	return s
}

func analyzeSignals(data DisplayData) int {
	newDigits := make(map[int]Digit)

	// get the easy ones
	for _, sig := range data.signals {
		switch len(sig) {
		case len(one):
			newDigits[1] = buildDigit(sig)
		case len(four):
			newDigits[4] = buildDigit(sig)
		case len(seven):
			newDigits[7] = buildDigit(sig)
		case len(eight):
			newDigits[8] = buildDigit(sig)
		}
	}

	// for _, sig := range data.signals {
	// 	switch len(sig) {
	// 	case 6:
	// 		// check for 0
	// 		temp := buildDigit(sig)

	// 		_, ok := newDigits[0]
	// 		if len(difference(temp, newDigits[7])) == 3 && !ok {
	// 			newDigits[0] = temp
	// 		}
	// 	}
	// }

	for _, sig := range data.signals {
		switch len(sig) {
		case 6:
			// check for 9
			temp := buildDigit(sig)
			s := difference(temp, newDigits[4])
			// fmt.Printf("s when looking for 9: %s\n", d2s(s))
			_, ok := newDigits[9]
			if len(s) == 2 && !ok {
				// fmt.Printf("found 9: %s\n", sig)
				// fmt.Printf("found 9: %s\n", d2s(temp))
				newDigits[9] = temp
			}
		}
	}

	// get the top segment 'a'
	aa := difference(newDigits[7], newDigits[1])
	// fmt.Printf("aa: %s\n", d2s(aa))

	// get bottom left segment 'e'
	ee := difference(newDigits[8], newDigits[9])
	// fmt.Printf("ee: %s\n", d2s(ee))

	for _, sig := range data.signals {
		switch len(sig) {
		case 6:
			// check for 0
			temp := buildDigit(sig)
			s := difference(difference(temp, newDigits[1]), union(aa, ee))

			// fmt.Printf("s when looking for 9: %s\n", d2s(s))
			_, ok := newDigits[0]
			if len(s) == 2 && !ok {
				// fmt.Printf("found 0: %s\n", sig)
				// fmt.Printf("found 0: %s\n", d2s(temp))
				newDigits[0] = temp
			}
		}
	}

	// get middle line 'd'
	dd := difference(newDigits[8], newDigits[0])
	// fmt.Printf("dd: %s\n", d2s(dd))

	// get bottom line 'g'
	gg := difference(newDigits[8], union(union(newDigits[4], ee), aa))
	// fmt.Printf("gg: %s\n", d2s(gg))

	// build 3 from knowns
	newDigits[3] = union(union(aa, union(dd, gg)), newDigits[1])

	// get top line 'b'
	bb := difference(newDigits[8], union(newDigits[3], ee))

	// fmt.Printf("3: %s, ee: %s, union: %s\n", d2s(newDigits[3]), d2s(ee), d2s(union(newDigits[3], ee)))
	// fmt.Printf("BB: %s\n", d2s(bb))

	for _, sig := range data.signals {
		switch len(sig) {
		case 5:
			// find 5 or 2
			temp := buildDigit(sig)
			if !reflect.DeepEqual(temp, newDigits[3]) {

				d1 := difference(newDigits[8], temp)
				d2 := difference(d1, bb)
				// fmt.Printf("number: %s, d1: %s, d2: %s, bb: %s\n", sig, d2s(d1), d2s(d2), d2s(bb))
				if len(d2) == 1 {
					newDigits[2] = temp
					// fmt.Printf("Found 2: %s\n", sig)
				} else {
					newDigits[5] = temp
					// fmt.Printf("Found 5: %s\n", sig)
				}
			}
		case 6:
			// find 6, since 9 and 0 should already be done
			temp := buildDigit(sig)
			if !reflect.DeepEqual(temp, newDigits[0]) && !reflect.DeepEqual(temp, newDigits[9]) {
				newDigits[6] = temp
			}
		}
	}

	factor := 1000
	res := 0
	for _, o := range data.outputs {
		output := buildDigit(o)
		for k, digit := range newDigits {
			if reflect.DeepEqual(output, digit) {
				res += k * factor
			}
		}
		factor /= 10
	}
	// if res < 1000 {
	// 	for k, v := range newDigits {
	// 		fmt.Printf("%s: %d\n", d2s(v), k)
	// 	}
	// }
	fmt.Printf("%v: %d\n", data.outputs, res)
	return res
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	data := parseInput(string(input))
	sum := 0
	for _, d := range data {
		sum += analyzeSignals(d)
	}
	fmt.Printf("total sum: %d\n", sum)
}
