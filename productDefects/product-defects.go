package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

const DefectiveProduct int32 = 1

/*
 * Complete the 'largestArea' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY samples as parameter.
 */
func largestArea(samples [][]int32) int32 {
	length := len(samples)
	dp := make([][]int32, length+1, length+1)
	var maxSqLen int32 = 0
	for i := 1; i <= length; i++ {
		dp[i-1] = initialize(dp[i-1], length+1)
		dp[i] = initialize(dp[i], length+1)
		for j := 1; j <= length; j++ {
			if samples[i-1][j-1] == DefectiveProduct {
				dp[i][j] = int32(math.Min(math.Min(float64(dp[i][j-1]), float64(dp[i-1][j])), float64(dp[i-1][j-1]))) + 1
				maxSqLen = int32(math.Max(float64(maxSqLen), float64(dp[i][j])))
			}
		}
	}
	return maxSqLen
}

func initialize(arr []int32, length int) []int32 {
	if arr == nil {
		arr = make([]int32, length, length)
	}
	return arr
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	samplesRows, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	samplesColumns, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var samples [][]int32
	for i := 0; i < int(samplesRows); i++ {
		samplesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var samplesRow []int32
		for _, samplesRowItem := range samplesRowTemp {
			samplesItemTemp, err := strconv.ParseInt(samplesRowItem, 10, 64)
			checkError(err)
			samplesItem := int32(samplesItemTemp)
			samplesRow = append(samplesRow, samplesItem)
		}

		if len(samplesRow) != int(samplesColumns) {
			panic("Bad input")
		}

		samples = append(samples, samplesRow)
	}

	result := largestArea(samples)

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
