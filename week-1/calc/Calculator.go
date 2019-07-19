package calc

import (
	"log"
	"strconv"
)

func Calculate(inputParams []string) (int, error) {
	sumParams := make([]int, 0)
	var temp int
	var err error
	var isNegative bool
	for i := 0; i < len(inputParams); i++ {
		param := inputParams[i]
		switch param {
		case "*":
			i++
			nextVal := convertStrToInt(inputParams[i])
			temp = multiply(temp, nextVal)
		case "/":
			i++
			nextVal := convertStrToInt(inputParams[i])
			temp, err = division(temp, nextVal)
			if err != nil {
				return 0, err
			}
		case "-":
			if isNegative {
				temp = multiply(temp, -1)
				isNegative = false
			}
			sumParams = append(sumParams, temp)
			isNegative = true
		case "+":
			if isNegative {
				temp = multiply(temp, -1)
				isNegative = false
			}
			sumParams = append(sumParams, temp)
		default:
			temp = convertStrToInt(param)
		}
	}
	if isNegative {
		temp = multiply(temp, -1)
		isNegative = false
	}
	// Add last value into sum array
	sumParams = append(sumParams, temp)
	return sum(sumParams...), err
}

func convertStrToInt(input string) int {
	result, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err.Error())
	}
	return result
}
