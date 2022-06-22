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
