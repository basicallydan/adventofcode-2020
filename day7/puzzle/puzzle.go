package puzzle

import (
  "regexp"
  "strings"
  // "math"
)

func CanHoldGoldBag(ruleString string) (string, string, bool) {
  var regex = *regexp.MustCompile(`([a-z]+ [a-z]+) bags contain (.*)\.`)
  var matches = regex.FindStringSubmatch(ruleString)

  var colour = matches[1]
  var rule = matches[2]

  return colour, rule, strings.Contains(rule, "shiny gold")
}

func CanHoldGoldBagGivenRules(ruleString string, existingRules []string) (string, bool) {
  var colour, _, direct = CanHoldGoldBag(ruleString)

  if direct == true {
    return colour, true
  }

  return "", false
}
