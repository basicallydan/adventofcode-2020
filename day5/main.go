package main

import(
	"fmt"
	"./puzzle"
	"../read_file"
)

func main() {
	var input = read_file.ReadFile("/Users/dan/Code/adventofcode-2020/day5/input.txt")
	// fmt.Println(puzzle.DetermineHighestSeatID(input))
	fmt.Println(puzzle.DetermineMissingSeatID(input))
}
