package main

import (
	"testing"
)

func calcValue(arr []int32) int64 {
	var oddsum int64
	var evensum int64
	for i := 0; i < len(arr); i++ {
		if i%2 == 0 {
			evensum += int64(arr[i])
		} else {
			oddsum += int64(arr[i])
		}
	}
	return (evensum - oddsum) * (evensum - oddsum)
}

// not optimized - timeout
func solutionRough(arr []int32) int32 {
	// Write your code here
	var maxOut int64
	// var size = len(arr)
	for len(arr) > 1 {
		value := calcValue(arr)
		if value > maxOut {
			maxOut = value
		}
		arr = arr[1:]
	}

	return int32(maxOut)
}

// optimized - a variant of  "finding a subarray of the maximum sum" algorithm
func solutionOptimized(arr []int32) int32 {
	var maxP int32 = 0
	var sum int32 = 0
	size := len(arr)
	for i := 0; i < size; i++ {
		val := arr[i]
		if i%2 == 0 {
			val = -val
		}

		if val <= sum+val {
			sum += val
		} else {
			sum = val
		}
		if sum > maxP {
			maxP = sum
		}
	}

	var maxN int32 = 0
	sum = 0
	for i := 0; i < size; i++ {
		val := arr[i]
		if i%2 == 0 {
			val = -val
		}

		if val >= sum+val {
			sum += val
		} else {
			sum = val
		}
		if sum < maxN {
			maxN = sum
		}

	}

	maxP = maxP * maxP
	maxN = maxN * maxN

	if maxP < maxN {
		return maxN
	}
	return maxP
}

func TestFindCount(t *testing.T) {

	cases := []struct {
		Args           []int32
		ExpectedOutput int32
	}{
		{[]int32{2, -1, -4, 5}, 81},
		{[]int32{-1, -4, 2}, 36},
	}

	t.Run("Test FindCount", func(t *testing.T) {
		for _, tc := range cases {
			var actualOutput int32 = solutionOptimized(tc.Args) //, len(tc.Args[0])-1)
			if tc.ExpectedOutput != actualOutput {
				t.Errorf("Wrong output for args: %v, expected: %v, got: %v",
					tc.Args, tc.ExpectedOutput, actualOutput)
			}
		}
	})

}
