package main

import "fmt"

func main() {
	amounts := []int{7500, 25000, 75000, -500, 0, 9999, 10000, 49999, 50000, -1}

	for _, amount := range amounts {
		printCategory(amount)
		fmt.Println()
	}
}

func printCategory(amount int) {
	if amount < 0 {
		fmt.Println("Ошибка: сумма заказа не может быть отрицательной")
		return
	}

	var category string
	switch {
	case amount < 10000:
		category = "small"
	case amount < 50000:
		category = "regular"
	default:
		category = "large"
	}

	fmt.Println(fmt.Sprintf("Сумма заказа: %d коп.", amount))
	fmt.Println(fmt.Sprintf("Категория: %s", category))
}
