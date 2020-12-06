package main

import(
	"fmt"
	"./puzzle"
	"../read_file"
)

func main() {
	var input = read_file.ReadFile("/Users/dan/Code/adventofcode-2020/day4/input.txt")
	fmt.Println(puzzle.NumberOfValidPassports(input))
}
