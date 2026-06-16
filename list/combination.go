package list

func Combinations[T any](lists [][]T) [][]T {
	combinations := [][]T{}
	for _, list := range lists {
		newCombinations := [][]T{}
		for _, item := range list {
			if len(combinations) == 0 {
				newCombinations = append(newCombinations, []T{item})
			} else {
				for _, combination := range combinations {
					newCombination := append(combination, item)
					newCombinations = append(newCombinations, newCombination)
				}
			}
		}
		combinations = newCombinations
	}

	return combinations
}
