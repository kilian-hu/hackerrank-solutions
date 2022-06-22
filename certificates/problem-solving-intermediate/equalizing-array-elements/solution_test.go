package main

import (
	"sort"
	"testing"
)

func solutionOptimized(args []int32) int32 {
	var size = len(args)
	if size < 3 {
		return -1
	}
	k, d := args[size-2], args[size-1]
	size -= 2
	if size < int(k) {
		return -1
	}
	a := args[:size]

	max := 200_000

	var v = make([][]int, max, max)
	var n = len(a)

	// Traverse the array
	for i := 0; i < n; i++ {
		var cnt = 0
		var val = a[i]

		v[val] = append(v[val], 0)

		for val > 0 {
			val /= d
			cnt++
			v[val] = append(v[val], cnt)
		}
	}

	var ans = max

	// minimum count of moves
	for i := 0; i < max; i++ {

		// Check if there are at least
		// K equal elements for v[i]
		if len(v[i]) >= int(k) {
			move := 0

			sort.Ints(v[i])

			// Add the sum of minimum K moves
			for j := 0; j < int(k); j++ {
				move += v[i][j]
			}

			// Update answer
			if ans > move {
				ans = move
			}
		}
	}

	// Return the final answer
	return int32(ans)
}

func TestSolution(t *testing.T) {

	cases := []struct {
		Args           []int32
		ExpectedOutput int32
	}{
		{[]int32{1, 2, 3, 4, 5, 3, 2}, 2},
		{[]int32{64, 30, 25, 33, 2, 2}, 3},
		{[]int32{1, 2, 3, 4, 4, 3}, 6},
		{[]int32{1, 2, 3, 4, 5, 6, 7, 8, 22, 11, 24, 13, 3}, 34},
		{[]int32{1, 2, 3, 4, 5, 6, 7, 8, 3333, 123211, 134234, 23121, 2323, 12421, 1212, 2323, 22, 333, 13213, 222, 15123, 1222, 3333, 11, 24, 13, 3}, 34},
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
