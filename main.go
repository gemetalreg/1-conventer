package main

import "fmt"

func main() {
	curr1 := userInput()
	num := userInputNum()
	curr2 := userInputTarget(curr1)

	m := map[string]map[string]float64{}

	m["USD"] = map[string]float64{}
	m["EUR"] = map[string]float64{}
	m["RUB"] = map[string]float64{}

	m["USD"]["EUR"] = 0.85
	m["USD"]["RUB"] = 76.23
	m["EUR"]["USD"] = 1.18
	m["EUR"]["RUB"] = 89.65
	m["RUB"]["USD"] = 0.013
	m["RUB"]["EUR"] = 0.011

	res := calc(&m, num, curr1, curr2)
	fmt.Println(res)

}

func userInput() string {
	var input string
	for {
		fmt.Print("Введите исходную валюту USD, EUR, RUB")
		fmt.Scan(&input)
		if input == "USD" || input == "EUR" || input == "RUB" {
			break
		}
	}
	return input
}

func userInputNum() float64 {
	var input float64
	for {
		fmt.Print("Введите число")
		_, err := fmt.Scan(&input)

		if err != nil {
			fmt.Println("Введите число")
		} else {
			break
		}
	}
	return input
}

func userInputTarget(curr1 string) string {
	var input string
	for {
		fmt.Print("Введите целевую валюту USD, EUR, RUB")
		fmt.Scan(&input)
		if input == curr1 {
			fmt.Printf("Ваша целевая валюта %s не должна совпадать с целевой валютой %s\n", curr1, input)
			continue
		}

		if input == "USD" || input == "EUR" || input == "RUB" {
			break
		}
	}
	return input
}

func calc(m *map[string]map[string]float64, count float64, curr1 string, curr2 string) (res float64) {

	return (*m)[curr1][curr2] * count
}
