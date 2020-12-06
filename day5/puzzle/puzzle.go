package puzzle

import (
  // "regexp"
  // "fmt"
	"sort"
  "strings"
  "math"
)

func DetermineSeatRowInRange(seatCode string, lowest int, highest int) int {
  // Go does floor division apparently!
  var rowRange = highest - lowest
  var rangeDiff = int(math.Ceil(float64(rowRange) / 2.0))
  if seatCode[0] == 'F' {
    return DetermineSeatRowInRange(seatCode[1:], lowest, highest - rangeDiff)
  }
  if seatCode[0] == 'B' {
    return DetermineSeatRowInRange(seatCode[1:], lowest + rangeDiff, highest)
  }
  return highest
}

func DetermineSeatRow(seatCode string) int {
  return DetermineSeatRowInRange(seatCode, 0, 127)
}

func DetermineSeatColumnInRange(seatCode string, lowest int, highest int) int {
  if len(seatCode) == 0 {
    return highest
  }

  var rowRange = highest - lowest
  var rangeDiff = int(math.Ceil(float64(rowRange) / 2.0))
  if seatCode[0] == 'L' {
    return DetermineSeatColumnInRange(seatCode[1:], lowest, highest - rangeDiff)
  }
  if seatCode[0] == 'R' {
    return DetermineSeatColumnInRange(seatCode[1:], lowest + rangeDiff, highest)
  }

  return DetermineSeatColumnInRange(seatCode[1:], lowest, highest)
}

func DetermineSeatColumn(seatCode string) int {
  return DetermineSeatColumnInRange(seatCode, 0, 7)
}

func DetermineSeatID(seatCode string) int {
  return DetermineSeatRow(seatCode) * 8 + DetermineSeatColumn(seatCode)
}

func DetermineHighestSeatID(seatCodes string) int {
  var codes = strings.Split(seatCodes, "\n")
  var highest = 0

	for _, seatCode := range codes {
    var newID = DetermineSeatID(seatCode)
    if newID > highest {
      highest = newID
    }
  }
  
  return highest
}

func DetermineEverySeatID(seatCodes string) []int {
  var codes = strings.Split(seatCodes, "\n")
  var everyID []int

	for _, seatCode := range codes {
    everyID = append(everyID, DetermineSeatID(seatCode))
  }

  sort.Ints(everyID)
  
  return everyID
}

func DetermineMissingSeatID(seatCodes string) int {
  var everyID = DetermineEverySeatID(seatCodes)

	for index, seatID := range everyID {
    if seatID == everyID[index + 1] - 2 {
      return seatID + 1
    }
  }


  
  return 0
}
