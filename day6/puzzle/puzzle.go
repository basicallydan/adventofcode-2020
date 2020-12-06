package puzzle

import (
  "regexp"
  "strings"
  // "math"
)

func CharsWithYesAnswers(groupString string) map[string]bool {
  var set = make(map[string]bool) // New empty set
  var regex = *regexp.MustCompile(`[a-z]`)
  var alphaChars = regex.FindAllString(groupString, -1)
  for _, c := range alphaChars {
    set[c] = true
  }
  return set
}

func NumberOfYesAnswersFromSet(charsWithYes map[string]bool) int {
  return len(charsWithYes)
}

func NumberOfYesAnswers(groupString string) int {
  var set = CharsWithYesAnswers(groupString)
  return len(set)
}

func SumOfNumberOfYesAnswers(groupsString string) int {
  var groups = strings.Split(groupsString, "\n\n")
  var sum = 0
  for _, groupString := range groups {
    sum += NumberOfYesAnswers(groupString)
  }
  return sum
}

func NumberOfUnanimousYesAnswers(groupString string) int {
  var charsWithYes = CharsWithYesAnswers(groupString)
  var totalUnanimous = NumberOfYesAnswersFromSet(charsWithYes)
  var peopleAnswers = strings.Split(groupString, "\n")

  for char, _ := range charsWithYes {
    for _, answer := range peopleAnswers {
      if !strings.Contains(answer, char) {
        totalUnanimous -= 1
        break
      }
    }
  }
  
  return totalUnanimous
}

func SumOfNumberOfUnanimousYesAnswers(groupsString string) int {
  var groups = strings.Split(groupsString, "\n\n")
  var sum = 0
  for _, groupString := range groups {
    sum += NumberOfUnanimousYesAnswers(groupString)
  }
  return sum
}
