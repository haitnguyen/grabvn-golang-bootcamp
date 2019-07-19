package main

import (
	"./calc"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Input expression > ")
	// input example: 1 + 2 * 3 + 4 / 2 - 5
	for scanner.Scan() {
		input := scanner.Text()
		input = strings.TrimSpace(input)
		expr := strings.Split(input, " ")
		result, err := calc.Calculate(expr)

		if err != nil {
			log.Fatal("calculate error: ", err.Error())
		} else {
			fmt.Println(input, " = ", result)
		}
		fmt.Println("Input next expression > ")

	}
}
