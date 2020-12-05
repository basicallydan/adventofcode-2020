package puzzle

import (
  // "regexp"
  // "fmt"
  "strings"
)

type passportType struct {
  byr int
  iyr int
  eyr int
  hgt_value int
  hgt_unit string
  hcl string
  ecl string
  pid string
}

func ParsePassport(passport string) passportType {
  return passportType{
    1980,
    2012,
    2030,
    74,
    "in",
    "#623a2f",
    "grn",
    "087499704",
  }
}

// M: byr (Birth Year)
// M: iyr (Issue Year)
// M: eyr (Expiration Year)
// M: hgt (Height)
// M: hcl (Hair Color)
// M: ecl (Eye Color)
// M: pid (Passport ID)
// O: cid (Country ID)
func PassportIsValid(passport string) bool {
  if !strings.Contains(passport, "byr") {
    return false
  }
  if !strings.Contains(passport, "iyr") {
    return false
  }
  if !strings.Contains(passport, "eyr") {
    return false
  }
  if !strings.Contains(passport, "hgt") {
    return false
  }
  if !strings.Contains(passport, "hcl") {
    return false
  }
  if !strings.Contains(passport, "ecl") {
    return false
  }
  if !strings.Contains(passport, "pid") {
    return false
  }
  return true
}

func NumberOfValidPassports(passports string) int {
  var passportArray = strings.Split(passports,"\n\n")
  var numberOfValidPassports = 0

	for _, passport := range passportArray {
    if PassportIsValid(passport) {
      numberOfValidPassports += 1
    }
	}

  return numberOfValidPassports
}
