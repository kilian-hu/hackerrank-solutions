package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'maxSubarrayValue' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */
// optimized
func maxSubarrayValue(arr []int32) int32 {
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

	result := maxSubarrayValue(arr)

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
