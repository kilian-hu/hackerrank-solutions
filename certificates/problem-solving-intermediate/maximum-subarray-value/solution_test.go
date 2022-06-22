package main

import (
	"testing"
)

// optimized - a variant of  "finding a subarray of the maximum sum" algorithm
func solutionOptimized(arr []int32) int32 {
	var max_p, max_n, sum_p, sum_n int32

	for i := 0; i < len(arr); i++ {
		val := arr[i]
		if i%2 != 0 {
			val = -val
		}
		// for positive
		if val <= sum_p+val {
			sum_p += val
		} else {
			sum_p = val
		}
		if sum_p > max_p {
			max_p = sum_p
		}
		// for negative
		if val >= sum_n+val {
			sum_n += val
		} else {
			sum_n = val
		}
		if sum_n < max_n {
			max_n = sum_n
		}
	}

	max_p = max_p * max_p
	max_n = max_n * max_n

	if max_p > max_n {
		return max_p
	}
	return max_n
}
func TestSolution(t *testing.T) {

	cases := []struct {
		Args           []int32
		ExpectedOutput int32
	}{
		{[]int32{2, -1, -4, 5}, 81},
		{[]int32{-1, -4, 2}, 36},
	}

	t.Run("Test Solution", func(t *testing.T) {
		for _, tc := range cases {
			var actualOutput int32 = solutionOptimized(tc.Args) //, len(tc.Args[0])-1)
			if tc.ExpectedOutput != actualOutput {
				t.Errorf("Wrong output for args: %v, expected: %v, got: %v",
					tc.Args, tc.ExpectedOutput, actualOutput)
			}
		}
	})

}
