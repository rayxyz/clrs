package main

import (
	"fmt"
)

func SayHi() {
	fmt.Println("Hi, you can call me Ray.")
}

func Sort(arr []int) {
	fmt.Println("Array to be sorted: ", arr)
	if len(arr) < 2 {
		fmt.Println("Length of Array less than 2. Program aborted.")
	}
	for j := 1; j < len(arr); j++ {
		key := arr[j]
		i := j - 1
		for i >= 0 && arr[i] > key {
			arr[i+1] = arr[i]
			i--
		}
		arr[i+1] = key
	}
	fmt.Println("Array after sorted: ", arr)
}
