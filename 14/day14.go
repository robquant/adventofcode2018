package main

import (
	"fmt"
	"strings"
)

func equal(a, b []int) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func main() {
	receipes := []int{3, 7}
	elves := []int{0, 1}

	nreceipes := 920831
	part1 := false
	target := []int{9, 2, 0, 8, 3, 1}
	for {
		s := receipes[elves[0]] + receipes[elves[1]]
		if s < 10 {
			receipes = append(receipes, s)
			if len(receipes) >= len(target) {
				if equal(target, receipes[len(receipes)-len(target):]) {
					fmt.Printf("Part2: %d\n", len(receipes)-len(target))
					break
				}
			}
		} else {
			receipes = append(receipes, s/10)
			receipes = append(receipes, s%10)
			if len(receipes) >= len(target) {
				if equal(target, receipes[len(receipes)-len(target):]) {
					fmt.Printf("Part2: %d\n", len(receipes)-len(target))
					break
				}
			}
			if len(receipes) >= len(target)+1 {
				if equal(target, receipes[len(receipes)-len(target)-1:len(receipes)-1]) {
					fmt.Printf("Part2: %d\n", len(receipes)-len(target)-1)
					break
				}
			}
		}
		elves[0] = (elves[0] + receipes[elves[0]] + 1) % len(receipes)
		elves[1] = (elves[1] + receipes[elves[1]] + 1) % len(receipes)
		if len(receipes) >= nreceipes+10 && !part1 {
			var writer strings.Builder
			for _, n := range receipes[nreceipes : nreceipes+10] {
				fmt.Fprintf(&writer, "%d", n)
			}
			fmt.Println(writer.String())
			part1 = true
		}
	}

}
