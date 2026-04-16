package main

import "fmt"

func main() {
	const UsdToEur = 0.85
	const UsdToRub = 75.38
	const EurToRub = UsdToRub / UsdToEur

	input := userInput()
}

func userInput() string {
	var input string

	fmt.Scan(&input)
	return input
}

func calc(count float64, curr1 string, curr2 string) (res float64) {

	return
}
