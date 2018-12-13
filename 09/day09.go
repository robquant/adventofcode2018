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

func main() {
	last_marble := 1618
	players := 10
	scores := make(map[int]int)
	current := &node{value: 0}
	current.next = current
	current.prev = current
	for stone := 1; stone <= last_marble; stone++ {
		player := stone % players
		if stone%23 == 0 {
			for i := 0; i < 6; i++ {
				current = current.prev
			}
			removed := remove(current)
			if val, ok := scores[stone]; ok {
				scores[player] = val + stone + removed.value
			} else {
				scores[player] = stone + removed.value
			}
		} else {
			node := &node{value: stone}
			insert(current.next, node)
			current = node
		}
	}
	fmt.Printf("%v\n", scores)
}
