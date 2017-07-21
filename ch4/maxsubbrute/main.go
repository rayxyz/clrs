package main

import (
	"fmt"
	"math"
)

func findMaxSubarryBrute(array []int) (maxBegin, maxEnd, sum int) {
	maxBegin, maxEnd, sum = 0, 0, math.MinInt64
	if array != nil && len(array) > 0 {
		sumtmp := 0
		for i := 0; i < len(array); i++ {
			sumtmp = array[i]
			for j := i + 1; j < len(array)-1; j++ {
				sumtmp += array[j]
				if sum < sumtmp {
					sum = sumtmp
					maxBegin = i
					maxEnd = j
				}
			}
		}
	}
	return maxBegin, maxEnd, sum
}

func main() {
	array := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	maxBegin, maxEnd, sum := findMaxSubarryBrute(array)
	fmt.Println("maxBegin => ", maxBegin, ", maxEnd => ", maxEnd, ", sum => ", sum)
}
