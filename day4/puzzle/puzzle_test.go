package puzzle

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestParsePassport_ValidButParseable(t *testing.T) {
  var input = "pid:087499704 hgt:74in\necl:grn iyr:2012\neyr:2030 byr:1980 hcl:#623a2f"
  var output = ParsePassport(input)
  assert.Equal(t, passportType{
    1980,
    2012,
    2030,
    74,
    "in",
    "#623a2f",
    "grn",
    "087499704",
  }, output)
}

func TestParsePassport_InvalidButParseable(t *testing.T) {
  var input = "eyr:1972 cid:100 hcl:#18171d ecl:amb hgt:170 pid:087499504 iyr:2018 byr:1926"
  var output = ParsePassport(input)
  assert.Equal(t, passportType{
    1926,
    2018,
    1972,
    186,
    "cm",
    "#18171d",
    "amb",
    "087499504",
  }, output)
}

func TestPassportIsValid_PassportIsInvalid(t *testing.T) {
  var input = "eyr:1972 cid:100 hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926"
  var output = PassportIsValid(input)
  assert.Equal(t, false, output)
}

func TestPassportIsValid_PassportIsInvalid2(t *testing.T) {
  var input = "iyr:2019 hcl:#602927 eyr:1967 hgt:170cm ecl:grn pid:012533040 byr:1946"
  var output = PassportIsValid(input)
  assert.Equal(t, false, output)
}

func TestPassportIsValid_PassportIsInvalid3(t *testing.T) {
  var input = "hcl:dab227 iyr:2012 ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277"
  var output = PassportIsValid(input)
  assert.Equal(t, false, output)
}

func TestPassportIsValid_PassportIsInvalid4(t *testing.T) {
  var input = "hgt:59cm ecl:zzz eyr:2038 hcl:74454a iyr:2023 pid:3556412378 byr:2007"
  var output = PassportIsValid(input)
  assert.Equal(t, false, output)
}

func TestPassportIsValid_PassportIsValid(t *testing.T) {
  var input = "pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980 hcl:#623a2f"
  var output = PassportIsValid(input)
  assert.Equal(t, true, output)
}

func TestPassportIsValid_PassportIsValid2(t *testing.T) {
  var input = "eyr:2029 ecl:blu cid:129 byr:1989 iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm"
  var output = PassportIsValid(input)
  assert.Equal(t, true, output)
}

func TestPassportIsValid_PassportIsValid3(t *testing.T) {
  var input = "hcl:#888785 hgt:164cm byr:2001 iyr:2015 cid:88 pid:545766238 ecl:hzl eyr:2022"
  var output = PassportIsValid(input)
  assert.Equal(t, true, output)
}

func TestPassportIsValid_PassportIsValid4(t *testing.T) {
  var input = "iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719"
  var output = PassportIsValid(input)
  assert.Equal(t, true, output)
}

func TestNumberOfValidPassports_Three(t *testing.T) {
  var input = "eyr:1972 cid:100 hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926" // Invalid
  input += "\n\niyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719" // Valid
  input += "\n\nhcl:#888785 hgt:164cm byr:2001 iyr:2015 cid:88 pid:545766238 ecl:hzl eyr:2022" // Valid
  input += "\n\nhgt:59cm ecl:zzz eyr:2038 hcl:74454a iyr:2023 pid:3556412378 byr:2007" // Invalid
  input += "\n\npid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980 hcl:#623a2f" // Valid
  var output = NumberOfValidPassports(input)
  assert.Equal(t, 3, output)
}
