package puzzle

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestParsePasswordItem(t *testing.T) {
  var input = "1-3 a: abcde"
  var output = ParsePasswordItem(input)
  assert.Equal(t, passwordPolicyTestItem{1, 3, "a", "abcde"}, output)
}

func TestNumberOfValidPasswordsFirstInput(t *testing.T) {
  var input = [3]string{"1-3 a: abcde","1-3 b: cdefg","2-9 c: ccccccccc"}
  var output = NumberOfValidPasswords(input[:])
  assert.Equal(t, 2, output)
}

func TestNumberOfValidPasswordsSecondInput(t *testing.T) {
  var input = [4]string{"2-5 a: abcdae","1-3 b: cdefg","2-9 b: ccccccccc","1-8 h: abcdefgij"}
  var output = NumberOfValidPasswords(input[:])
  assert.Equal(t, 1, output)
}

func TestNewPolicyNumberOfValidPasswordsFirstInput(t *testing.T) {
  var input = [3]string{"1-3 a: abcde","1-3 b: cdefg","2-9 d: dcccccccc"}
  var output = NewPolicyNumberOfValidPasswords(input[:])
  assert.Equal(t, 1, output)
}

func TestNewPolicyNumberOfValidPasswordsSecondInput(t *testing.T) {
  var input = [5]string{"6-14 c: cpcccgcddfzcgcccc","11-14 m: mmmmmmmmmmmmmmmkmmm","2-4 h: vhthh","1-8 h: hbcdefgij","1-8 a: baaaaaab"}
  var output = NewPolicyNumberOfValidPasswords(input[:])
  assert.Equal(t, 2, output)
}
