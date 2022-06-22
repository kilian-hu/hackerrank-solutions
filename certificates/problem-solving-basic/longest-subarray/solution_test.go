package main

import (
	"fmt"
	"testing"
)

func Abs(val int32) int32 {
	if val > 0 {
		return val
	} else {
		return -val
	}
}

func solve(arr []int32) int32 {
	// Write your code here
	var max int
	var scope [][]int32

	for len(arr) > 0 {
		var elements []int32
		var lastIdx int
		var subArray []int32
		elements = append(elements, arr[0])
		subArray = append(subArray, arr[0])
		// fmt.Println(arr)
		for i := 1; i < len(arr); i++ {
			if arr[i] == arr[0] && len(elements) == 1 {
				lastIdx = i
			}
			if arr[i] != elements[len(elements)-1] && arr[i] != elements[0] {
				elements = append(elements, arr[i])
			}
			if len(elements) > 2 || Abs(arr[i]-arr[i-1]) > 1 {
				break
			}
			subArray = append(subArray, arr[i])
		}
		if len(subArray) > max {
			max = len(subArray)
		}
		// fmt.Println(lastIdx)
		arr = arr[lastIdx+1:]
		scope = append(scope, subArray)

	}
	fmt.Println(scope)
	return int32(max)
}

func TestFindCount(t *testing.T) {

	cases := []struct {
		Args           []int32
		ExpectedOutput int32
	}{
		{[]int32{0, 1, 2, 1, 2, 3}, 4},
		{[]int32{1, 1, 1, 3, 3, 2, 2}, 4},
	}

	t.Run("Test FindCount", func(t *testing.T) {
		for _, tc := range cases {
			var actualOutput int32 = solve(tc.Args) //, len(tc.Args[0])-1)
			if tc.ExpectedOutput != actualOutput {
				t.Errorf("Wrong output for args: %v, expected: %v, got: %v",
					tc.Args, tc.ExpectedOutput, actualOutput)
			}
		}
	})

}
