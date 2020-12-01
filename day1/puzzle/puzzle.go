package puzzle

func GetExpenseOutputPart1(expenses []int) int {
  for _, number_one := range expenses {
    for _, number_two := range expenses {
      if (number_one + number_two == 2020) {
        return number_one * number_two
      }
    }
  }

	return 0
}

func GetExpenseOutputPart2(expenses []int) int {
  for _, number_one := range expenses {
    for _, number_two := range expenses {
      for _, number_three := range expenses {
        if (number_one + number_two + number_three == 2020) {
          return number_one * number_two * number_three
        }
      }
    }
  }

	return 0
}
