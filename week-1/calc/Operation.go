package calc

import "errors"

func sum(values ...int) int {
	result := 0
	for _, val := range values {
		result += val
	}
	return result
}

func subtract(values ...int) int {
	result := 0
	for index, val := range values {
		if index == 0 {
			result = val
		} else {
			result -= val
		}
	}
	return result
}

func multiply(values ...int) int {
	result := 1
	for _, val := range values {
		if val == 0 {
			return 0
		}
		result *= val
	}
	return result
}

func division(values ...int) (int, error) {
	var result int
	for index, val := range values {
		if index == 0 && val == 0 {
			return 0, nil
		} else if index == 0 && val != 0 {
			result = val
		} else if index != 0 && val == 0 {
			return 0, errors.New("can not divide zero")
		} else {
			result /= val
		}
	}
	return result, nil
}
