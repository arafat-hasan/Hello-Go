package main

import "fmt"

func Sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, value := range nums {
		total += value
	}
	fmt.Println(total)
}

func main() {
	Sum(1, 2)
	Sum(1, 2, 3)

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	Sum(nums...)
}
