package main

import (
	"bufio"
	"fmt"
	"io"
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
	prevTop, prevBottom, prevCount := 0, 0, -1
	var response int32 = 1
	for _, row := range samples {
		top, bottom, count := findRange(row)
		if (prevCount >= count || prevCount == -1) && top == prevTop {
			response = int32(count)
		}
		if prevTop != top && prevBottom != bottom {
			response = 1
		}
		prevTop = top
		prevCount = count

	}
	return response

}

func findRange(row []int32) (top int, bottom int, count int) {
	top = -1
	bottom = -1

	var prev int32 = -1
	for i, item := range row {
		if item == DefectiveProduct {
			if prev == DefectiveProduct || prev == -1 {
				count++
				bottom = i
			}
			if top == -1 || prev != DefectiveProduct {
				top = i
			}
		}
		prev = item
	}

	return top, bottom, count
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
