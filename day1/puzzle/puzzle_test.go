package puzzle

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestGetExpenseOutputOne(t *testing.T) {
  var input = [4]int{2000, 20, 40, 1000}
  var output = GetExpenseOutputPart1(input[:])
  assert.Equal(t, output, 40000)
}

func TestGetExpenseOutputTwo(t *testing.T) {
  var input = [6]int{1721,979,366,299,675,1456}
  var output = GetExpenseOutputPart1(input[:])
  assert.Equal(t, output, 514579)
}
