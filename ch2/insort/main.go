package main

import (
	"fmt"
	// "clrs/insort"
)

func main() {
	fmt.Println("I'm in.")
	// insort.SayHi()
	// insort.Sort([]int{23, 32, 1, 4, 45, 27})
	// LinearSearch([]int{23, 32, 1, 4, 45, 27}, 5475)
	SelectionSort([]int{23, 32, 1, 4, 45, 27})
	LinuxFunc()
}

func LinuxFunc() {
	fmt.Println("I'm in Linux. Fuck, yea!!!")
}

func SelectionSort(A []int) {
	fmt.Println("Befor selection sort: ", A)
	for i := 0; i < len(A)-1; i++ {
		k := i
		tmp := A[i]
		for j := i + 1; j < len(A); j++ {
			if A[i] > A[j] {
				tmp = A[i]
				A[i] = A[j]
				A[j] = tmp
			}
		}
		A[k] = A[i]
	}
	fmt.Println("After selection sort: ", A)
}

func LinearSearch(arr []int, value int) {
	var index int
	var exists bool = false
	for i := 0; i < len(arr); i++ {
		if value == arr[i] {
			index = i
			exists = true
			break
		}
	}
	if exists {
		fmt.Println("The index is: ", index)
	} else {
		fmt.Println("The value ", value, " doesn't exist in array ", arr)
	}

}
