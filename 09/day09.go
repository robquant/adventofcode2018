package main

import "fmt"

// clock-wise == direction of next
// counter clock-wise == direction of prev

type node struct {
	next, prev *node
	value      int
}

//insert after (clock_wise) the current
func insert(current, node *node) {
	next := current.next
	current.next = node
	node.next = next
	node.prev = current
	next.prev = node
}

// remove prev node (counter clock-wise)
func remove(current *node) *node {
	removed := current.prev
	current.prev = removed.prev
	current.prev.next = current
	return removed
}

func print(start *node) {
	run := start
	fmt.Printf("%d->", run.value)
	run = run.next
	for ; run != start; run = run.next {
		fmt.Printf("%d->", run.value)
	}
	fmt.Printf("\n")
}

func highscore(scores map[int]int) int {
	score := 0
	for _, v := range scores {
		if v > score {
			score = v
		}
	}
	return score
}

func play(lastMarble, players int) int {
	scores := make(map[int]int)
	current := &node{value: 0}
	current.next = current
	current.prev = current
	for stone := 1; stone <= lastMarble; stone++ {
		player := stone % players
		if stone%23 == 0 {
			for i := 0; i < 6; i++ {
				current = current.prev
			}
			removed := remove(current)
			if val, ok := scores[player]; ok {
				scores[player] = val + stone + removed.value
			} else {
				scores[player] = stone + removed.value
			}
		} else {
			node := &node{value: stone}
			insert(current.next, node)
			current = node
		}
		// print(current)
	}
	return highscore(scores)
}

func main() {
	fmt.Printf("%v\n", play(71240, 478))
	fmt.Printf("%v\n", play(71240*100, 478))
}
