package main

import "fmt"

func Sum(numbers []int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	fmt.Println(Sum(numbers))
	fmt.Println(Sum([]int{}))
	fmt.Println(Sum(nil))
	fmt.Println(Sum([]int{-2, -1, 0, 1, 2}))

	numbers = append(numbers, 12)
	fmt.Println(len(numbers), Sum(numbers))

	shared := numbers
	shared[0] = 100
	fmt.Println(numbers[0], shared[0])
}
