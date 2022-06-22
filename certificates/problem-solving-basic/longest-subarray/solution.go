package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func Abs(val int32) int32 {
	if val > 0 {
		return val
	} else {
		return -val
	}
}

/*
 * Complete the 'longestSubarray' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */
func longestSubarray(arr []int32) int32 {
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

	result := longestSubarray(arr)

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
