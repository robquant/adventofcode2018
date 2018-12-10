package main

import (
	"fmt"
)

func sum(numbers []int) int {
	res := 0
	for _, n := range numbers {
		res += n
	}
	return res
}

func part1(numbers []int) {
	fmt.Println(sum(numbers))
}

func part2(numbers []int) {
	freqs := make(map[int]bool)
	freq := 0
	for {
		for _, n := range numbers {
			freq += n
			if ok, _ := freqs[freq]; ok {
				fmt.Println(freq)
				return
			}
			freqs[freq] = true
		}
	}
}

func main() {
	numbers := read()
	part1(numbers)
	part2(numbers)
}
