package puzzle

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func Test_DetermineSeatRow_FBFBBFFRLR_44(t *testing.T) {
  var input = "FBFBBFFRLR"
  var output = DetermineSeatRow(input)
  assert.Equal(t, 44, output)
}

func Test_DetermineSeatRow_BFFFBBFRRR_70(t *testing.T) {
  var input = "BFFFBBFRRR"
  var output = DetermineSeatRow(input)
  assert.Equal(t, 70, output)
}

func Test_DetermineSeatColumn_BFFFBBFRRR_7(t *testing.T) {
  var input = "BFFFBBFRRR"
  var output = DetermineSeatColumn(input)
  assert.Equal(t, 7, output)
}

func Test_DetermineSeatID_BFFFBBFRRR_567(t *testing.T) {
  var input = "BFFFBBFRRR"
  var output = DetermineSeatID(input)
  assert.Equal(t, 567, output)
}

func Test_DetermineSeatID_BBFFBBFRRR_567(t *testing.T) {
  var input = "BBFFBBFRRR"
  var output = DetermineSeatID(input)
  assert.Equal(t, 823, output)
}

func Test_DetermineSeatID_BBBFBBFRRR_567(t *testing.T) {
  var input = "BBBFBBFRRR"
  var output = DetermineSeatID(input)
  assert.Equal(t, 951, output)
}

func Test_DetermineSeatRow_FFFBBBFRRR_14(t *testing.T) {
  var input = "FFFBBBFRRR"
  var output = DetermineSeatRow(input)
  assert.Equal(t, 14, output)
}

func Test_DetermineSeatColumn_FFFBBBFRRR_7(t *testing.T) {
  var input = "FFFBBBFRRR"
  var output = DetermineSeatColumn(input)
  assert.Equal(t, 7, output)
}

func Test_DetermineSeatRow_BBFFBBFRLL_102(t *testing.T) {
  var input = "BBFFBBFRLL"
  var output = DetermineSeatRow(input)
  assert.Equal(t, 102, output)
}

func Test_DetermineSeatColumn_BBFFBBFRLL_4(t *testing.T) {
  var input = "BBFFBBFRLL"
  var output = DetermineSeatColumn(input)
  assert.Equal(t, 4, output)
}

func Test_DetermineHighestSeatID(t *testing.T) {
  var input = "BFFFBBFRRR\nBBBFBBFRRR\nBBFFBBFRRR"
  var output = DetermineHighestSeatID(input)
  assert.Equal(t, 951, output)
}
