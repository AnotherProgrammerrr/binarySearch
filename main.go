package main

import "fmt"

func main() {
	target := 7
	array := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	left := 0 
	right := len(array) - 1


	for left <= right {
		mid := left + (right - left) / 2

		if array[mid] == target {
			fmt.Printf("target %d at index %d\n", target, mid)
			return
		} else if array[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
			
		}
	}
}