package puzzle

import (
  // "regexp"
  "fmt"
  "strings"
)

func NumberOfTrees(treeMap string, xDiff int, yDiff int) int {
  var numTrees = 0
  var rows = strings.Split(treeMap, "\n")
  var width = len(rows[0])
  var height = len(rows)
  fmt.Printf("yDiff is %d\n", yDiff)
  fmt.Printf("xDiff is %d\n", xDiff)
  fmt.Printf("Height is %d\n", height)
  fmt.Printf("Width is %d\n", width)

  var x = 0
  var y = 0

  for y < height {
    fmt.Printf("At x = %d, y = %d it is", x, y)
    var atPosition = strings.Split(rows[y], "")[x]
    fmt.Printf(" %s\n", atPosition)
    if atPosition == "#" {
      numTrees += 1
    }

    x += xDiff

    if x >= width {
      x -= width
    }

    y += yDiff
    fmt.Printf("EOL X = %d, Y = %d\n", x, y)
  }

  return numTrees
}
