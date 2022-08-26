package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for index, el := range nums {
		element, isIn := m[target-el]

		if isIn && (el != target-el || index != element) {
			return []int{index, element}
		} else {
			m[el] = index
		}
	}

	return []int{0, 0}
}
