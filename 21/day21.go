package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type registers [6]int
type opcode func(*registers, int, int, int)

func addr(reg *registers, a, b, c int) {
	reg[c] = reg[a] + reg[b]
}

func addi(reg *registers, a, b, c int) {
	reg[c] = reg[a] + b
}

func mulr(reg *registers, a, b, c int) {
	reg[c] = reg[a] * reg[b]
}

func muli(reg *registers, a, b, c int) {
	reg[c] = reg[a] * b
}

func banr(reg *registers, a, b, c int) {
	reg[c] = reg[a] & reg[b]
}

func bani(reg *registers, a, b, c int) {
	reg[c] = reg[a] & b
}

func borr(reg *registers, a, b, c int) {
	reg[c] = reg[a] | reg[b]
}

func bori(reg *registers, a, b, c int) {
	reg[c] = reg[a] | b
}

func setr(reg *registers, a, b, c int) {
	reg[c] = reg[a]
}

func seti(reg *registers, a, b, c int) {
	reg[c] = a
}

func gtir(reg *registers, a, b, c int) {
	if a > reg[b] {
		reg[c] = 1
	} else {
		reg[c] = 0
	}
}

func gtri(reg *registers, a, b, c int) {
	if reg[a] > b {
		reg[c] = 1
	} else {
		reg[c] = 0
	}
}

func gtrr(reg *registers, a, b, c int) {
	if reg[a] > reg[b] {
		reg[c] = 1
	} else {
		reg[c] = 0
	}
}

func eqir(reg *registers, a, b, c int) {
	if a == reg[b] {
		reg[c] = 1
	} else {
		reg[c] = 0
	}
}

func eqri(reg *registers, a, b, c int) {
	if reg[a] == b {
		reg[c] = 1
	} else {
		reg[c] = 0
	}
}

func eqrr(reg *registers, a, b, c int) {
	if reg[a] == reg[b] {
		reg[c] = 1
	} else {
		reg[c] = 0
	}
}

type Instruction struct {
	opcode opcode
	params [3]int
}

func readProgram(filename string, opcodes map[string]opcode) ([]Instruction, int) {
	instructions := make([]Instruction, 0)
	ins := Instruction{}
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	ip := 0
	for _, line := range strings.Split(string(input), "\n") {
		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "#") {
			items := strings.Split(line, " ")
			ip, _ = strconv.Atoi(items[1])
			continue
		}
		for i, item := range strings.Split(line, " ") {
			if i == 0 {
				ins.opcode = opcodes[item]
			} else {
				number, err := strconv.Atoi(strings.TrimSpace(item))
				if err != nil {
					log.Fatal(err)
				}
				ins.params[i-1] = number
			}
		}
		instructions = append(instructions, ins)
	}
	return instructions, ip
}

func run(program []Instruction, ipReg, start int, printStats bool) registers {
	regs := registers{start, 0, 0, 0, 0, 0}
	ip := 0
	counter := 0
	stats := make(map[int]int)
	for i := 0; i < len(program); i++ {
		stats[i] = 0
	}
	for ; ip < len(program); ip++ {
		if printStats && counter%1000000 == 0 {
			for i := 0; i < len(program); i++ {
				fmt.Printf("%d: %f\n", i, float64(stats[i])/float64(counter)*100)
			}
		}
		stats[ip]++
		ins := program[ip]
		regs[ipReg] = ip
		ins.opcode(&regs, ins.params[0], ins.params[1], ins.params[2])
		ip = regs[ipReg]
		counter++
	}
	return regs
}

func main() {

	opcodes := map[string]opcode{
		"addr": addr,
		"addi": addi,
		"mulr": mulr,
		"muli": muli,
		"banr": banr,
		"bani": bani,
		"bori": bori,
		"borr": borr,
		"setr": setr,
		"seti": seti,
		"gtir": gtir,
		"gtri": gtri,
		"gtrr": gtrr,
		"eqir": eqir,
		"eqri": eqri,
		"eqrr": eqrr,
	}

	program, ip := readProgram("input.txt", opcodes)
	fmt.Printf("%v\n", run(program, ip, 0, false))
}
