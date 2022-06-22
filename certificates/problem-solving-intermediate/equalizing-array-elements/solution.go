package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func minOperations(args []int32) int32 {
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

			//fmt.Println(i, v[i], move)

			// Update answer
			if ans > move {
				ans = move
			}
		}
	}

	// Return the final answer
	return int32(ans)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	arrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var arr []int32

	for i := 0; i < int(arrCount); i++ {
		arrItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := minOperations(arr)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
