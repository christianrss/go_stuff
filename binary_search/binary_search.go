package main

import "fmt"

func binary_search(array []int64, to_search int64) bool {
	found := false
	low := 0
	high := len(array) - 1
	for low <= high {
		mid := (low + high) / 2
		if array[mid] == to_search {
			found = true
			break
		}
		if array[mid] > to_search {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return found
}

func main() {
	//arg1 := os.Args[1]

	if binary_search([]int64{1, 2, 3, 4}, 2) {
		fmt.Println("Found")
	} else {
		fmt.Println("Not Found")
	}
}
