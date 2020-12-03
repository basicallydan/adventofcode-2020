package main

import(
	"fmt"
	"./puzzle"
)

func main() {
	var input = ""
	input += "......#...........#...#........\n"
	input += ".#.....#...##.......#.....##...\n"
	input += "......#.#....#.................\n"
	input += "..............#.#.......#......\n"
	input += ".....#.#...##...#.#..#..#..#..#\n"
	input += ".......##...#..#...........#...\n"
	input += ".......#.##.#...#.#.........#..\n"
	input += "..#...##............##......#.#\n"
	input += ".......#.......##......##.##.#.\n"
	input += "...#...#........#....#........#\n"
	input += "#............###.#......#.....#\n"
	input += "..#........#....#..#..........#\n"
	input += "..#..##....#......#..#......#..\n"
	input += "........#......#......#..#..#..\n"
	input += "..#...#....#..##.......#.#.....\n"
	input += ".....#.#......#..#....#.##.#..#\n"
	input += "......###.....#..#..........#..\n"
	input += ".#................#.#..........\n"
	input += ".........#..#...#......##......\n"
	input += "##...#....#...#.#...#.##..#....\n"
	input += "...##...#....#.........###.....\n"
	input += ".#.#....#.........##...........\n"
	input += "....#.#..#..#...........#......\n"
	input += "..#..#.#....#....#...#.........\n"
	input += "..........##.....#.##..........\n"
	input += "..#.#....#..##......#.#.....##.\n"
	input += "..#...#.##......#..........#...\n"
	input += "......#....#..#.....#.....#...#\n"
	input += "#.#...##.#.##.........#..#.....\n"
	input += "...#.#.#.........#.....#.#.#...\n"
	input += "..#.........#...............#..\n"
	input += "#..##.....#.........#....#.....\n"
	input += "...#....##..##...........##..#.\n"
	input += "......##.................#.#...\n"
	input += "##.......#....#.#.#.....#......\n"
	input += "....#.#...#.................##.\n"
	input += "#...#.........##.....#.........\n"
	input += "#....#.###..#.....##.#....#....\n"
	input += "#..#....#...#....#.#.#.........\n"
	input += ".......#...........#....#.....#\n"
	input += "#...#.............#........#...\n"
	input += ".......#.....#...#..#.........#\n"
	input += ".##.....##.....##.......#......\n"
	input += "....##...##.......#..#.#.....#.\n"
	input += ".##.........#......#........##.\n"
	input += ".......#...#...###.#..#........\n"
	input += "..#..###......##..##...........\n"
	input += ".#..#......##..#.#.........#...\n"
	input += "...#.......#........#...#.#....\n"
	input += "...#....#..#....#.....##.......\n"
	input += "............#......#..........#\n"
	input += ".#.......#......#.#....#..#.#..\n"
	input += "##.........#.#.#..........#....\n"
	input += "....##.....#...................\n"
	input += ".......#..#........#...........\n"
	input += "....##.#..#......###.......#...\n"
	input += "....#....#...#.#......#...#...#\n"
	input += ".......#.....##..#....#...#....\n"
	input += "#...#........#.........#..##...\n"
	input += "...........##.........#.#...#..\n"
	input += "....................#....#.##..\n"
	input += ".#..#..#.........#....#..#..##.\n"
	input += "......................#........\n"
	input += "..###....#.......#.....###.##..\n"
	input += "......#......#.......#.....#..#\n"
	input += ".....#...#.##...#......#....#..\n"
	input += ".....#.....##.............#....\n"
	input += "....#......##..#....#.......#..\n"
	input += ".##....#..##......###....#..#..\n"
	input += "...###.#.............##...#.#..\n"
	input += ".....#.....#.....#...#..#.#....\n"
	input += "..#.#.....###......#.......#...\n"
	input += "..........#.##......#.........#\n"
	input += "..##..#.......................#\n"
	input += "........#......#............#..\n"
	input += "#..#..#..#.#......#..#....#....\n"
	input += "...##......#.............#....#\n"
	input += "...........#..#..##.......#....\n"
	input += ".....#.........#.#..#..........\n"
	input += "##...#.......#.#....#..#..#....\n"
	input += "#.#.#...........#.##.#.#..###..\n"
	input += "#..#...........#.........##....\n"
	input += "............#.#..............#.\n"
	input += ".#....#....##.#...........#..#.\n"
	input += "....#...#..#...#....#....#.....\n"
	input += "....#....#...#..#......#.......\n"
	input += ".#.#.........#.......#.##......\n"
	input += ".#..##...#........#...........#\n"
	input += "##...#..#...#...#.....#...#....\n"
	input += "....###.#..#.......##.#..#...#.\n"
	input += "...##.......####...##.#........\n"
	input += "#....#....#.#............#..#..\n"
	input += "#.#.#...#...................##.\n"
	input += "##......#...........#..........\n"
	input += "#..#..#....#.#...#......#......\n"
	input += ".##...#.....#...#........#.....\n"
	input += "..#............#..............#\n"
	input += "###........#..#....#...#......#\n"
	input += "###..##......#.##...........#..\n"
	input += "........#......#..#.....#......\n"
	input += "...#..........#..#...........#.\n"
	input += "....#..#..#....#........#....#.\n"
	input += ".#.................#####..##..#\n"
	input += ".....#...##..#..........#.##...\n"
	input += "..#..............#...####......\n"
	input += ".....#.##..................#.#.\n"
	input += "...#.#..#..#........#..........\n"
	input += "...........#....#.#..#.........\n"
	input += ".....##.......#......#..#.#.#..\n"
	input += "...#.............##...#........\n"
	input += "...............#.......##.##.##\n"
	input += ".....#........#........#.#..#..\n"
	input += "...#..#.........#...##...###...\n"
	input += "...#.#.............###.#.....#.\n"
	input += ".#..........#......###.#.#.....\n"
	input += "....##..##.............###.....\n"
	input += "..#..#.#...##...#.......##.....\n"
	input += "..........###........#.....#.#.\n"
	input += "#.#....#..#..#......#...#...#..\n"
	input += ".........#......##.......#.#..#\n"
	input += "...#.....#.........##..#..#....\n"
	input += ".....##.#..##.##..##...........\n"
	input += "...#.#.##....#..#..#......#..#.\n"
	input += "#....#....#.............#...##.\n"
	input += "#......#..#.####.#.##.#....##..\n"
	input += "##.#.#....##..................#\n"
	input += ".....##......#.......##.......#\n"
	input += "..#......#.#..#...##......##...\n"
	input += "..#....##....#.........#..##...\n"
	input += ".###.....#....##...........#...\n"
	input += ".........#......#.#........#...\n"
	input += "...#...#..#.#....######.#..#...\n"
	input += "###......#.#.#.........##.#....\n"
	input += ".....#...#.........#...#.......\n"
	input += "....#.............#.#.........#\n"
	input += "..##...#...#.......#......#....\n"
	input += ".....#...#.#...#...#..#........\n"
	input += ".#......#......................\n"
	input += "...###..#..#....#...##.#.......\n"
	input += ".#.#.....##...#...#.....#...##.\n"
	input += ".....###..###....##............\n"
	input += ".....##....#..#.....#.##.......\n"
	input += "#........#.........#...#..#....\n"
	input += "...#.#.........#..#.......#.#..\n"
	input += "....#.#....##.....#..........#.\n"
	input += ".#..#....#..#.#..#..#.........#\n"
	input += "#...#....#..............#......\n"
	input += ".........#.....#.##...##...###.\n"
	input += ".....#....##............#..#...\n"
	input += ".....#.#...........#..#....#...\n"
	input += ".#..........#...#......#.....#.\n"
	input += ".#...........#.....#..#........\n"
	input += "..............#......##...#..#.\n"
	input += "...#.........#..#....#..##...##\n"
	input += "..##...#..................#....\n"
	input += "#.....#.................#......\n"
	input += "...#......#..#..........#.#....\n"
	input += "......#..#.....#.....##...#..#.\n"
	input += "......#........#..........#....\n"
	input += "...##.##....#..##.#..........#.\n"
	input += "..........#..#.#.##............\n"
	input += "..##........................#..\n"
	input += ".....#.#.#......#....#....##...\n"
	input += "#....#.........#........#......\n"
	input += ".##.......#...#...#........##..\n"
	input += "....##......#....#.#..........#\n"
	input += "..#.......#..............#.....\n"
	input += ".....#......#.#...#..#.#.#....#\n"
	input += ".....#..#........#.##.##.......\n"
	input += "##........#..........#.........\n"
	input += ".....#..##....#.#......###..##.\n"
	input += "#.#...##.........#.#.....#..#..\n"
	input += "#....#.#...#........#.....#..#.\n"
	input += "........................#......\n"
	input += "....###......#............#...#\n"
	input += "...#..##......#..##.........#..\n"
	input += ".............#...#......#..#..#\n"
	input += "....#......#....#...........#..\n"
	input += "..#.#.####.#.....##........#..#\n"
	input += "#..#...#..#..#.......#.#..#....\n"
	input += "..#..#..#....#.#.........##..#.\n"
	input += ".......#......#.#............#.\n"
	input += "...#.............#.#.....#.....\n"
	input += "...#.#.........##...#.#.......#\n"
	input += "........#...#...........##...#.\n"
	input += "..........#....#......#....##..\n"
	input += "..........#...........#........\n"
	input += "...#..#...#..........#......#..\n"
	input += "......#......#....#.....#..#.#.\n"
	input += "........##.................#..#\n"
	input += ".#........#.#...........#......\n"
	input += "#...#........#.#.#.....#.#.#...\n"
	input += ".........#........#..#..#....#.\n"
	input += "##........#..........#....#..#.\n"
	input += ".#.##...........#..#.#..##....#\n"
	input += ".......#.#....#..#......#......\n"
	input += "..#.....#........##..#......###\n"
	input += "..#...#..................#....#\n"
	input += "......#...#..#.##.......#......\n"
	input += "........#...#.#................\n"
	input += ".........#............#........\n"
	input += "..#.....##....#.#..##..........\n"
	input += "#.....#..........#....#........\n"
	input += "....#.#...#...##....#.....##...\n"
	input += "..#.#.......#.............#...#\n"
	input += "...##..............#......#....\n"
	input += "#......#...#................##.\n"
	input += ".#.#...#.#..#..................\n"
	input += "...##.......#...........#.#.#..\n"
	input += "#......#.#.#........#.##...####\n"
	input += ".......#..#.#.........#.#.##..#\n"
	input += "..............#....#.........#.\n"
	input += "...........#.#..#....##......#.\n"
	input += "#.............#...##..#.......#\n"
	input += ".........#............#...#.##.\n"
	input += ".......#.........#.#.....#..#..\n"
	input += "........................#.#.##.\n"
	input += "#......#.#......#.........#....\n"
	input += "...#.......#.......#.....#.....\n"
	input += "#..#....#................#...#.\n"
	input += "........#.#..##......#.........\n"
	input += "#..#...##....##....##.........#\n"
	input += ".......#...#...###.............\n"
	input += "#.#..#........#.#.#............\n"
	input += "#.....#........##.........#.#..\n"
	input += ".#..........#....#.#....###....\n"
	input += ".#.....#...#.#........#..#.##..\n"
	input += "...#.##......#..#.............#\n"
	input += "..##..#.#...................#..\n"
	input += ".....#....#...#.#...#...#......\n"
	input += ".....#..#.#....#.#.............\n"
	input += "#.#....#.#.##..###..........#..\n"
	input += "........#.#.............#..#...\n"
	input += ".........#.......#.............\n"
	input += ".##.#............##...#........\n"
	input += "......#................#.......\n"
	input += "...............#..#...........#\n"
	input += "...#.......#...#.##.....#....#.\n"
	input += "##..##..#..........#...........\n"
	input += ".##.#.......#...#..#...#...#...\n"
	input += "....#..#...........#....#.##...\n"
	input += ".#........#........#....#......\n"
	input += ".......#...#.##.#..#.#..#......\n"
	input += ".#..#......#....#...##....#.#..\n"
	input += "......#...##.#.....##.###.....#\n"
	input += ".#....#..#......#...#.#.....#..\n"
	input += "#............#....##...##.##...\n"
	input += "#...#.#....#...#.......##...##.\n"
	input += "#...........#.##..#....#.....#.\n"
	input += "...#..#...#.........#.......#..\n"
	input += ".#....#.....#............#.#..#\n"
	input += ".#.....#.#...#.#....##......###\n"
	input += "..#..#.#.#...#..#.............#\n"
	input += "...#...#..#....#........#...##.\n"
	input += ".......#.....#...##...........#\n"
	input += "#.##.................#...##...#\n"
	input += "..............##........#.....#\n"
	input += "............#...#..#.......#.#.\n"
	input += "#.#.....#.........#...#......#.\n"
	input += "#.###..#......#..#..#...#.....#\n"
	input += ".....#.......#.................\n"
	input += "........#..#......#.#...#......\n"
	input += "#.......#..#........#...#..#...\n"
	input += "..#...#.......##.............#.\n"
	input += "#.......#.......##...#.........\n"
	input += ".........#....#.#..##.....#...#\n"
	input += "..#.....#.#.......#....#.......\n"
	input += "...#.......#.....#..##.#..#....\n"
	input += "....#.......#.#.#..............\n"
	input += ".#..#......#........#.#..##..##\n"
	input += "....#...#.##.#...#....##...#...\n"
	input += "#..##..#.....#.......#.........\n"
	input += "....#..#..#.#............#.....\n"
	input += "#.......##...##..##............\n"
	input += "...............................\n"
	input += "....#.......#.##...#.....#.#...\n"
	input += "...#........#....#.#..#..#.....\n"
	input += "##.......#.....##.#.#....#....#\n"
	input += "#.............#...........#.##.\n"
	input += "#...........#.#..........#.....\n"
	input += "#..#....#....#.#.........#.#...\n"
	input += "......#.#.#..#.#.#.............\n"
	input += "...#.....#........##....#......\n"
	input += "..#...#...#.#.......#......#...\n"
	input += ".##........#...#..#..........#.\n"
	input += "..#...........#..##.....##.....\n"
	input += "............#..#.#...#.....#...\n"
	input += "..........#....##.......#......\n"
	input += "....#....#.................#..#\n"
	input += "....#...............#.........#\n"
	input += "..#.#...#......#..........##...\n"
	input += ".....#...........#.........#..#\n"
	input += ".......#.....##.....#.#........\n"
	input += ".#.#..........#....#...........\n"
	input += ".#..##....#........#....#......\n"
	input += "....#.#..#.......#..#.........#\n"
	input += "..#....#.....#......#..#.......\n"
	input += "......#........#.......#...#.#.\n"
	input += ".......#.......#....#.....##...\n"
	input += "....##........#..#...#.#..#...#\n"
	input += ".#......#...........##....#....\n"
	input += "##....##......#.......#.......#\n"
	input += ".##....#.##......#.......##..#.\n"
	input += "...#..#.#.#.......#..#.###.....\n"
	input += "..........##....#..#.##........\n"
	input += "...#........###.#..#........#..\n"
	input += ".....#....#..##....#.....#....#\n"
	input += "#..........#..........#.#....#.\n"
	input += "..#....#.....#..............#..\n"
	input += "#..................#......#.##.\n"
	input += ".#...#.#.....#.........##......\n"
	input += "...#...........#.....#......#..\n"
	input += "......#.....#.#..##......##....\n"
	input += "...#....###..#.....#..#..##..##\n"
	input += "......#.......##..#..#.........\n"
	input += "#..#.#....#.#..#..........##.#.\n"
	input += "..#..#..##..#.#.#.#.....#......\n"
	input += "..#.#...#..#.....###.#.........\n"
	input += "##.#.#......#........#.####....\n"
	input += ".............#..#..#....#......\n"
	input += "...##..........#.......#.#....#\n"
	input += "..#.....................#......\n"
	input += "..#..#...##...#.##........#...."
	var slopeProduct = puzzle.NumberOfTrees(input, 1, 1)
	slopeProduct = slopeProduct * puzzle.NumberOfTrees(input, 3, 1)
	slopeProduct = slopeProduct * puzzle.NumberOfTrees(input, 5, 1)
	slopeProduct = slopeProduct * puzzle.NumberOfTrees(input, 7, 1)
	slopeProduct = slopeProduct * puzzle.NumberOfTrees(input, 1, 2)
	fmt.Println(slopeProduct)
}
