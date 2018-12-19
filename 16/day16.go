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

type Instruction struct {
	opcode int
	params [3]int
}

type Sample struct {
	input, output [4]int
	Instruction
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
				if i == 0 {
					sample.opcode = number
				} else {
					sample.params[i-1] = number
				}
			}
		}
	}
	return samples
}

func readProgram(filename string) []Instruction {
	instructions := make([]Instruction, 0)
	ins := Instruction{}
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range strings.Split(string(input), "\n") {
		if len(line) == 0 {
			continue
		}
		for i, numberstr := range strings.Split(line, " ") {
			number, err := strconv.Atoi(strings.TrimSpace(numberstr))
			if err != nil {
				log.Fatal(err)
			}
			if i == 0 {
				ins.opcode = number
			} else {
				ins.params[i-1] = number
			}
		}
		instructions = append(instructions, ins)
	}
	return instructions
}

func run(program []Instruction, opcodeToFuncIndex map[int]int, opcodes []opcode) [4]int {
	regs := [4]int{}
	for _, ins := range program {
		op := opcodes[opcodeToFuncIndex[ins.opcode]]
		op(&regs, ins.params[0], ins.params[1], ins.params[2])
	}
	return regs
}

func union(a, b []int) []int {
	res := make([]int, 0)
	for _, vala := range a {
		for _, valb := range b {
			if vala == valb {
				res = append(res, vala)
			}
		}
	}
	return res
}

func collaps(opcodeToFuncIndex map[int][]int) map[int]int {
	res := make(map[int]int)
	for {
		single := -1
		for k, v := range opcodeToFuncIndex {
			if len(v) == 1 {
				res[k] = v[0]
				single = v[0]
				delete(opcodeToFuncIndex, k)
				break
			}
		}
		if single >= 0 {
			for k, v := range opcodeToFuncIndex {
				// Remove single from v
				found := -1
				for i, val := range v {
					if val == single {
						found = i
						break
					}
				}
				if found >= 0 {
					opcodeToFuncIndex[k] = append(v[:found], v[found+1:]...)
				}
			}
		} else {
			return res
		}
	}
}

func main() {

	opcodes := []opcode{addr, addi, mulr, muli, banr, bani, bori, borr, setr, seti, gtir, gtri, gtrr, eqir, eqri, eqrr}

	samples := readInput("input.txt")
	fmt.Printf("Number of samples: %d\n", len(samples))
	moreThanThreeInterp := 0
	opcodeToFuncIndexList := make(map[int][]int)
	allOpcodes := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	for i := 0; i < 16; i++ {
		opcodeToFuncIndexList[i] = append([]int{}, allOpcodes...)
	}
	for _, sample := range samples {
		validOpcodes := make([]int, 0)
		for opcodeIndex, opcode := range opcodes {
			var input [4]int
			copy(input[:], sample.input[:])
			opcode(&input, sample.params[0], sample.params[1], sample.params[2])
			if equal(sample.output, input) {
				validOpcodes = append(validOpcodes, opcodeIndex)
			}
		}
		if len(validOpcodes) >= 3 {
			moreThanThreeInterp++
		}
		opcodeToFuncIndexList[sample.opcode] = union(opcodeToFuncIndexList[sample.opcode], validOpcodes)
	}
	fmt.Printf("%d\n", moreThanThreeInterp)
	opcodeToFuncIndex := collaps(opcodeToFuncIndexList)
	program := readProgram("input2.txt")
	finalRegs := run(program, opcodeToFuncIndex, opcodes)
	fmt.Printf("%v\n", finalRegs)
}
