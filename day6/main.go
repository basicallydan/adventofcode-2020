package main

import(
	"fmt"
	"./puzzle"
	"../read_file"
)

func main() {
	var input = read_file.ReadFile("/Users/dan/Code/adventofcode-2020/day6/input.txt")
	// fmt.Println(puzzle.SumOfNumberOfYesAnswers(input))
	fmt.Println(puzzle.SumOfNumberOfUnanimousYesAnswers(input))
}
