package main

import (
	"fmt"
)

func main() {
	slice := []int{4, 2, 3, 2, 5, 4, 9}
	f, err := findMostOftenRepeated(slice)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)
	fmt.Println(slice)
}

func findMostOftenRepeated(array []int) (most int, err error) {
	if len(array) == 0 {
		return 0, fmt.Errorf("could not found max negative in empty slice")
	}
	myMap := make(map[int]int)
	maxCount := 0
	maxValue := 0

	// sly := []int{}
	for _, number := range array {
		myMap[number]++
	}
	for key, value := range myMap {
		if value >= maxCount {
			maxCount = value
			maxValue = key
		}
	}

	return maxValue, nil
}
