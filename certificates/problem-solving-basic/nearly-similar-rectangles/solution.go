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
 * Complete the 'nearlySimilarRectangles' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts 2D_LONG_INTEGER_ARRAY sides as parameter.
 */

func sumN(n int64) int64 {
	var sum int64 = 0
	var i int64
	for i = 1; i < n; i++ {
		sum += i
	}
	fmt.Println(n)
	fmt.Println(sum)
	return sum
}

func nearlySimilarRectangles(sides [][]int64) int64 {
	// Write your code here
	size := len(sides)
	if size < 2 {
		return 0
	}
	var ratios []float64

	for i := 0; i < size; i++ {
		ratio := float64(sides[i][0]) / float64(sides[i][1])
		ratios = append(ratios, ratio)
	}
	if size == 2 {
		if ratios[0] == ratios[1] {
			return 2
		} else {
			return 0
		}
	}

	var cnt int64
	for len(ratios) > 1 {
		var cnt_inner int64
		for i := 1; i < len(ratios); i++ {
			if ratios[0] == ratios[i] {
				cnt_inner++
				ratios = append(ratios[:i], ratios[i+1:]...)
				i--
			}
		}
		fmt.Println("cnti", cnt_inner)
		cnt += sumN(cnt_inner + 1)
		fmt.Println("cnt", cnt)
		ratios = ratios[1:]
	}
	return cnt
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	sidesRows, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	sidesColumns, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var sides [][]int64
	for i := 0; i < int(sidesRows); i++ {
		sidesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var sidesRow []int64
		for _, sidesRowItem := range sidesRowTemp {
			sidesItem, err := strconv.ParseInt(sidesRowItem, 10, 64)
			checkError(err)
			sidesRow = append(sidesRow, sidesItem)
		}

		if len(sidesRow) != int(sidesColumns) {
			panic("Bad input")
		}

		sides = append(sides, sidesRow)
	}

	result := nearlySimilarRectangles(sides)

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
