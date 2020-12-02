package puzzle

import (
  "regexp"
  "strconv"
)

type passwordPolicyTestItem struct {
  min int
  max int
  character string
  password string
}

func ParsePasswordItem(passwordItem string) passwordPolicyTestItem {
  
  // Looks like 2-6 z: zrzshvzlzkxzp
  
  regex := *regexp.MustCompile(`(\d+?)\-(\d+?) (\w{1}): (.+$)`)
  res := regex.FindStringSubmatch(passwordItem)


  min, _ := strconv.Atoi(res[1])
  max, _ := strconv.Atoi(res[2])
  character := res[3]
  password := res[4]

  item := passwordPolicyTestItem{min, max, character, password}

  return item
}

func NumberOfValidPasswords(policiesAndPasswords []string) int {
  numValid := 0
  for _, policyAndPassword := range policiesAndPasswords {
    parsedItem := ParsePasswordItem(policyAndPassword)
    regexString := parsedItem.character
    regex := *regexp.MustCompile(regexString)
    numberOfMatches := len(regex.FindAllString(parsedItem.password, -1))
    if (numberOfMatches >= parsedItem.min && numberOfMatches <= parsedItem.max) {
      numValid += 1
    }
  }
	return numValid
}

func NewPolicyNumberOfValidPasswords(policiesAndPasswords []string) int {
  numValid := 0
  for _, policyAndPassword := range policiesAndPasswords {
    numMatches := 0
    parsedItem := ParsePasswordItem(policyAndPassword)
    regexString := parsedItem.character
    regex := *regexp.MustCompile(regexString)
    allIndexes := regex.FindAllStringIndex(parsedItem.password, -1)
    for _, v := range allIndexes {
      if v[1] == parsedItem.min || v[1] == parsedItem.max {
        numMatches += 1
      }

      if (numMatches > 1) {
        break
      }
    }

    if (numMatches == 1) {
      numValid += 1
    }
  }
	return numValid
}
