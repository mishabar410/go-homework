package main

import "fmt"

func main() {
	orders := []int{12000, 7500, 31000, 4500}
	printStats(orders)

	fmt.Println()
	printStats([]int{9900})
	fmt.Println()
	printStats([]int{500, 500, 500})
	fmt.Println()
	printStats([]int{300, 100, 200})
	fmt.Println()
	printStats([]int{100, 200, 200})
	fmt.Println()
	printStats([]int{})
}

func printStats(orders []int) {
	if len(orders) == 0 {
		fmt.Println("Нет заказов для анализа")
		return
	}

	minOrder := orders[0]
	maxOrder := orders[0]
	total := 0

	for _, amount := range orders {
		if amount < minOrder {
			minOrder = amount
		}
		if amount > maxOrder {
			maxOrder = amount
		}
		total += amount
	}

	var average float64 = float64(total) / float64(len(orders))

	fmt.Println(fmt.Sprintf("Количество заказов: %d", len(orders)))
	fmt.Println(fmt.Sprintf("Минимальный заказ: %d коп.", minOrder))
	fmt.Println(fmt.Sprintf("Максимальный заказ: %d коп.", maxOrder))
	fmt.Println(fmt.Sprintf("Общая сумма: %d коп.", total))
	fmt.Println(fmt.Sprintf("Средний заказ: %g коп.", average))
}
