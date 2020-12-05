package main

import(
	"fmt"
	"./puzzle"
	"../read_file"
)

func main() {
	var input = read_file.ReadFile("/Users/danhough/Code/adventofcode2020/day4/input.txt")
	fmt.Println(puzzle.NumberOfValidPassports(input))
}
