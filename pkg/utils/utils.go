package utils

import "fmt"

func DeleteElementFromArray(arr []int, element int) {
	// Find the index of the element in the array
	indexToDelete := -1
	for i, val := range arr {
		if val == element {
			indexToDelete = i
			break
		}
	}

	// If the element is found, remove it from the array
	if indexToDelete != -1 {
		arr = append(arr[:indexToDelete], arr[indexToDelete+1:]...)
	} else {
		fmt.Println("Element not found in the array")
	}
}
