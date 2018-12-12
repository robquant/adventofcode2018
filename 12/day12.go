package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

const inital = "###....#..#..#......####.#..##..#..###......##.##..#...#.##.###.##.###.....#.###..#.#.##.#..#.#"
const EMPTY = '.'
const ALIVE = '#'

type rulesMap map[[5]byte]byte

func readRules(fname string) rulesMap {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rules := make(map[[5]byte]byte)
	for scanner.Scan() {
		fields := bytes.Split(scanner.Bytes(), []byte{' '})
		var key [5]byte
		copy(key[:], fields[0][:])
		rules[key] = []byte(fields[2])[0]
	}
	return rules
}

func getKey(state []byte, index int) [5]byte {
	key := [5]byte{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY}
	for i := index - 2; i <= index+2; i++ {
		if i < 0 || i >= len(state) {
			continue
		}
		key[i-index+2] = state[i]
	}
	return key
}

func generation(state []byte, rules rulesMap) []byte {
	result := make([]byte, 0)
	for i := -2; i <= len(state)+2; i++ {
		key := getKey(state, i)
		result = append(result, rules[key])
	}
	return result
}

func print(state []byte) {

	var builder strings.Builder
	for _, b := range state {
		builder.WriteByte(b)
	}
	fmt.Println(builder.String())
}

func sum(state []byte, steps int) int {
	sum := 0
	for i, pot := range state {
		if pot == ALIVE {
			sum += i - 2*steps
		}
	}
	return sum
}

const ALL_STEPS = 50000000000

func main() {
	state := []byte(inital)
	rules := readRules("rules")
	steps := 20
	for i := 1; i <= steps; i++ {
		state = generation(state, rules)
	}

	fmt.Printf("Summe: %d\n", sum(state, steps))

	state = []byte(inital)
	steps = 1000
	for i := 1; i <= steps; i++ {
		state = generation(state, rules)
	}
	sum1000 := sum(state, steps)
	fmt.Printf("Steps: %d Summe: %d\n", steps, sum1000)
	for i := 1; i <= steps; i++ {
		state = generation(state, rules)
	}
	sum2000 := sum(state, 2*steps)
	fmt.Printf("Steps: %d Summe: %d\n", 2*steps, sum2000)

	total := sum2000 + (ALL_STEPS-2*steps)/1000*(sum2000-sum1000)
	fmt.Printf("After %d steps: %d\n", ALL_STEPS, total)
}
