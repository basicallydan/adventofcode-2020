package puzzle

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func Test_NumberOfYesAnswers_TwoPeople(t *testing.T) {
  var input = "abc\nabx"
  var output = NumberOfYesAnswers(input)
  assert.Equal(t, 4, output)
}

func Test_NumberOfYesAnswers_FourPeople(t *testing.T) {
  var input = "a\nb\nc\nc"
  var output = NumberOfYesAnswers(input)
  assert.Equal(t, 3, output)
}

func Test_SumOfNumberOfYesAnswers_TwoGroups(t *testing.T) {
  var input = "abc\nabx\n\na\nb\nc\nc"
  var output = SumOfNumberOfYesAnswers(input)
  assert.Equal(t, 7, output)
}

func Test_NumberOfUnanimousYesAnswers_TwoPeople(t *testing.T) {
  var input = "abc\nabx"
  var output = NumberOfUnanimousYesAnswers(input)
  assert.Equal(t, 2, output)
}

func Test_NumberOfUnanimousYesAnswers_FourPeople(t *testing.T) {
  var input = "a\nb\nc\nc"
  var output = NumberOfUnanimousYesAnswers(input)
  assert.Equal(t, 0, output)
}

func Test_NumberOfUnanimousYesAnswers_FivePeople(t *testing.T) {
  var input = "abcd\nabce\nabcf\nabcf\nabcf"
  var output = NumberOfUnanimousYesAnswers(input)
  assert.Equal(t, 3, output)
}

func Test_SumOfNumberOfUnanimousYesAnswers_TwoGroups(t *testing.T) {
  var input = "abc\nabx\n\na\nb\nc\nc\n\nabcd\nabce\nabcf\nabcf\nabcf"
  var output = SumOfNumberOfUnanimousYesAnswers(input)
  assert.Equal(t, 5, output)
}
