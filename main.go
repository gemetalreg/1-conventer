package main

import "fmt"

func main() {
	curr1 := userInput()
	num := userInputNum()
	curr2 := userInputTarget(curr1)

	res := calc(num, curr1, curr2)
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

func calc(count float64, curr1 string, curr2 string) (res float64) {
	switch curr1 {
	case "USD":
		switch curr2 {
		case "EUR":
			return 0.85 * count
		case "RUB":
			return 76.23 * count
		}
	case "EUR":
		switch curr2 {
		case "USD":
			return 1.18 * count
		case "RUB":
			return 89.65 * count
		}
	case "RUB":
		switch curr2 {
		case "USD":
			return 0.013 * count
		case "EUR":
			return 0.011 * count
		}
	}
	return 0
}
