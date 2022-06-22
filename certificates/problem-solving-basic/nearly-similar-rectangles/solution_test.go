package main

import (
	"testing"
)

func solutionOptimized(sides [][]int64) int64 {
	// Write your code here
	size := len(sides)
	if size < 2 {
		return 0
	}
	// var ratios []float64
	var ratios = make(map[float64]int64)

	for i := 0; i < size; i++ {
		ratio := float64(sides[i][0]) / float64(sides[i][1])
		ratios[ratio]++
		// ratios = append(ratios, ratio)
	}
	var cnt int64
	for _, v := range ratios {
		cnt += v * (v - 1) / 2
	}

	return cnt
}

func TestSolution(t *testing.T) {

	cases := []struct {
		Args           [][]int64
		ExpectedOutput int64
	}{
		{[][]int64{[]int64{5, 10}, []int64{10, 10}, []int64{3, 6}, []int64{9, 9}}, 2},
		{[][]int64{[]int64{4, 8}, []int64{15, 30}, []int64{25, 50}}, 3},
		{[][]int64{[]int64{4, 8}, []int64{10, 20}, []int64{15, 30}, []int64{25, 50}}, 6},
		{[][]int64{[]int64{4, 8}, []int64{15, 30}}, 1},
	}

	t.Run("Test Solution", func(t *testing.T) {
		for _, tc := range cases {
			var actualOutput int64 = solutionOptimized(tc.Args) //, len(tc.Args[0])-1)
			if tc.ExpectedOutput != actualOutput {
				t.Errorf("Wrong output for args: %v, expected: %v, got: %v",
					tc.Args, tc.ExpectedOutput, actualOutput)
			}
		}
	})

}
