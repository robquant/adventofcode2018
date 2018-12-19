package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type opcode func(*[4]int, int, int, int)

func addr(reg *[4]int, a, b, c int) {
	reg[c] = reg[a] + reg[b]
}

func addi(reg *[4]int, a, b, c int) {
	reg[c] = reg[a] + b
}

func mulr(reg *[4]int, a, b, c int) {
	reg[c] = reg[a] * reg[b]
}

func muli(reg *[4]int, a, b, c int) {
	reg[c] = reg[a] * b
}

func banr(reg *[4]int, a, b, c int) {
	reg[c] = reg[a] & reg[b]
}

func bani(reg *[4]int, a, b, c int) {
	reg[c] = reg[a] & b
}

func borr(reg *[4]int, a, b, c int) {
	reg[c] = reg[a] | reg[b]
}

func bori(reg *[4]int, a, b, c int) {
	reg[c] = reg[a] | b
}

func setr(reg *[4]int, a, b, c int) {
	reg[c] = reg[a]
}

func seti(reg *[4]int, a, b, c int) {
	reg[c] = a
}

func gtir(reg *[4]int, a, b, c int) {
	if a > reg[b] {
		reg[c] = 1
	} else {
		reg[c] = 0
	}
}

func gtri(reg *[4]int, a, b, c int) {
	if reg[a] > b {
		reg[c] = 1
	} else {
		reg[c] = 0
	}
}

func gtrr(reg *[4]int, a, b, c int) {
	if reg[a] > reg[b] {
		reg[c] = 1
	} else {
		reg[c] = 0
	}
}

func eqir(reg *[4]int, a, b, c int) {
	if a == reg[b] {
		reg[c] = 1
	} else {
		reg[c] = 0
	}
}

func eqri(reg *[4]int, a, b, c int) {
	if reg[a] == b {
		reg[c] = 1
	} else {
		reg[c] = 0
	}
}

func eqrr(reg *[4]int, a, b, c int) {
	if reg[a] == reg[b] {
		reg[c] = 1
	} else {
		reg[c] = 0
	}
}

func equal(a, b [4]int) bool {
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

type Sample struct {
	input, output, opcodes [4]int
}

func readInput(filename string) []Sample {
	samples := make([]Sample, 0)
	sample := Sample{}
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range strings.Split(string(input), "\n") {
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, "Before") {
			numbers := line[strings.Index(line, "[")+1 : strings.Index(line, "]")]
			for i, numberstr := range strings.Split(numbers, ",") {
				number, err := strconv.Atoi(strings.TrimSpace(numberstr))
				if err != nil {
					log.Fatal(err)
				}
				sample.input[i] = number
			}
			continue
		}
		if strings.Contains(line, "After") {
			numbers := line[strings.Index(line, "[")+1 : strings.Index(line, "]")]
			for i, numberstr := range strings.Split(numbers, ",") {
				number, err := strconv.Atoi(strings.TrimSpace(numberstr))
				if err != nil {
					log.Fatal(err)
				}
				sample.output[i] = number
			}
			samples = append(samples, sample)
			continue
		} else {
			for i, numberstr := range strings.Split(line, " ") {
				number, err := strconv.Atoi(strings.TrimSpace(numberstr))
				if err != nil {
					log.Fatal(err)
				}
				sample.opcodes[i] = number
			}
		}
	}
	return samples
}

func main() {

	opcodes := []opcode{addr, addi, mulr, muli, banr, bani, bori, borr, setr, seti, gtir, gtri, gtrr, eqir, eqri, eqrr}

	samples := readInput("input.txt")
	fmt.Printf("Number of samples: %d\n", len(samples))
	moreThanThreeInterp := 0
	for _, sample := range samples {
		validOpcodes := make([]opcode, 0)
		for _, opcode := range opcodes {
			var input [4]int
			copy(input[:], sample.input[:])
			opcode(&input, sample.opcodes[1], sample.opcodes[2], sample.opcodes[3])
			if equal(sample.output, input) {
				validOpcodes = append(validOpcodes, opcode)
			}
		}
		if len(validOpcodes) >= 3 {
			moreThanThreeInterp++
		}
	}
	fmt.Printf("%d\n", moreThanThreeInterp)

}
