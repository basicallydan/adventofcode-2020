package puzzle

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestNumberOfTrees(t *testing.T) {
  var input =  "...#\n"
  input     += "..#.\n"
  input     += "#..."
  var output = NumberOfTrees(input, 2, 1)
  assert.Equal(t, 2, output)
}

func TestNumberOfTreesAOCExample(t *testing.T) {
  var input = ""
  input += "..##.......\n"
  input += "#...#...#..\n"
  input += ".#....#..#.\n"
  input += "..#.#...#.#\n"
  input += ".#...##..#.\n"
  input += "..#.##.....\n"
  input += ".#.#.#....#\n"
  input += ".#........#\n"
  input += "#.##...#...\n"
  input += "#...##....#\n"
  input += ".#..#...#.#"
  var output = NumberOfTrees(input, 3, 1)
  assert.Equal(t, 7, output)
}
