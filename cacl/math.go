package cacl

import "errors"

func Add(numbers ...int) (int, error) {
	sum := 0

	if len(numbers) < 2 {
		return sum, errors.New("provide more than 2 numbers")
	}
	for _, num := range numbers {
		sum += num
	}
	return sum, nil
}
