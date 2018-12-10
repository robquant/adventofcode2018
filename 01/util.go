package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"strconv"
)

func read() []int {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	numbers := make([]int, 0)
	for _, line := range bytes.Split(data, []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		number, err := strconv.Atoi(string(line))
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	return numbers
}
