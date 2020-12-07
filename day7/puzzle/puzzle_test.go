package puzzle

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func Test_CanHoldGoldBag_Direct_White_True(t *testing.T) {
  var input = "bright white bags contain 1 shiny gold bag."
  var colour, rules, output = CanHoldGoldBag(input)
  assert.Equal(t, "bright white", colour)
  assert.Equal(t, "1 shiny gold bag", rules)
  assert.Equal(t, true, output)
}

func Test_CanHoldGoldBag_Direct_Yellow_True(t *testing.T) {
  var input = "muted yellow bags contain 1 shiny gold bag."
  var colour, rules, output = CanHoldGoldBag(input)
  assert.Equal(t, "muted yellow", colour)
  assert.Equal(t, "1 shiny gold bag", rules)
  assert.Equal(t, true, output)
}

func Test_CanHoldGoldBag_Direct_Blue_False(t *testing.T) {
  var input = "pale blue bags contain 1 muted yellow bag."
  var colour, rules, output = CanHoldGoldBag(input)
  assert.Equal(t, "pale blue", colour)
  assert.Equal(t, "1 muted yellow bag", rules)
  assert.Equal(t, false, output)
}

func Test_CanHoldGoldBag_Indirect_White_True(t *testing.T) {
  var input = "light red bags contain 1 bright white bag, 2 muted yellow bags."
  var colour, output = CanHoldGoldBagGivenRules(input, []string{"bright white"})
  assert.Equal(t, "bright white", colour)
  assert.Equal(t, true, output)
}

func Test_CanHoldGoldBag_Indirect_Blue_True(t *testing.T) {
  var input = "pale blue bags contain 1 muted yellow bag, 1 shiny gold bag."
  var colour, output = CanHoldGoldBagGivenRules(input, []string{"bright white"})
  assert.Equal(t, "pale blue", colour)
  assert.Equal(t, true, output)
}

func Test_CanHoldGoldBag_Indirect_Yellow_True(t *testing.T) {
  var input = "light red bags contain 1 bright white bag, 2 muted yellow bags."
  var colour, output = CanHoldGoldBagGivenRules(input, []string{"muted yellow"})
  assert.Equal(t, "muted yellow", colour)
  assert.Equal(t, true, output)
}

func Test_CanHoldGoldBag_Indirect_Olive_False(t *testing.T) {
  var input = "light red bags contain 1 bright white bag, 2 muted yellow bags."
  var colour, output = CanHoldGoldBagGivenRules(input, []string{"dark olive"})
  assert.Equal(t, "dark olive", colour)
  assert.Equal(t, false, output)
}
