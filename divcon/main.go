/*
 mplementation of find the max subarray algorithm with
 divide and conquer method.
 @Ray
*/

package main

import (
	"fmt"
	"log"
	"math"
	_ "net/http/pprof"
)

func findMaxCrossingSubArray(a []int, low, mid, high int) (maxLeft, maxRight, sum int) {
	sum = 0
	// init leftSum and rightSum as negative infinity
	// and positive infinity, or just use math.Inf()
	// function instead.
	leftSum := math.MinInt32
	rightSum := math.MinInt32
	log.Println("max crossing: a => ", a)
	log.Println("max crossing: maxLeft => ", maxLeft, ", maxRight => ", maxRight, ", mid => ", mid)
	for i := mid; i > low; i-- {
		sum += a[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}
	sum = 0
	for j := mid + 1; j <= high; j++ {
		sum += a[j]
		if sum > rightSum {
			rightSum = sum
			maxRight = j
		}
	}
	return maxLeft, maxRight, (leftSum + rightSum)
}

func findMaxSubArray(a []int, low int, high int) (maxLeft, maxRight, sum int) {
	if high == low {
		return low, high, a[low]
	}
	mid := (low + high) / 2
	log.Println("max sub array: a => ", a)
	log.Println("max sub array: mid => ", mid, ", low => ", low, ", high => ", high)
	// m for max, l for left, and s for sum
	mll, mrl, sl := findMaxSubArray(a, low, mid)
	mlr, mrr, sr := findMaxSubArray(a, mid+1, high)
	mlc, mrc, sc := findMaxCrossingSubArray(a, low, mid, high)
	log.Println("max left: mll => ", mll, ", mrl => ", mrl, ", sl => ", sl)
	log.Println("max right: mlr => ", mlr, ", mrr => ", mrr, ", sr => ", sr)
	log.Println("max crossing: mlc => ", mlc, ", mrc => ", mrc, ", sc => ", sc)
	if sl >= sr && sl >= sc {
		return mll, mrl, sl
	}
	if sr > sl && sr >= sc {
		return mlr, mrr, sr
	}
	return mlc, mrc, sc
}

func main() {
	array := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	maxLeft, maxRight, sum := findMaxSubArray(array, 0, len(array)-1)
	// maxLeft, maxRight, sum := findMaxCrossingSubArray(array, 0, len(array)-1)
	fmt.Println("maxLeft: ", maxLeft, ", maxRight: ", maxRight, ", sum: ", sum)
}
