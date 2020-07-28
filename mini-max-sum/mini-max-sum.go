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

/*
Complete the miniMaxSum function below.
https://www.hackerrank.com/challenges/mini-max-sum/problem?h_r=next-challenge&h_v=zen&h_r=next-challenge&h_v=zen&h_r=next-challenge&h_v=zen
*/
func miniMaxSum(arr []int32) {
	min, max := miniMaxSumCalculator(arr)
	fmt.Printf("%d %d\n", min, max)
}

func miniMaxSumCalculator(arr []int32) (min int, max int) {
	for i := 0; i < len(arr); i++ {
		a := sum(arr[0:i])
		b := 0.0
		if i < len(arr) {
			b = sum(arr[i+1:])
		}
		c := a + b
		if min == 0 {
			min = int(c)
		} else {
			min = int(math.Min(c, float64(min)))
		}
		max = int(math.Max(c, float64(max)))
	}
	return min, max
}

func sum(arr []int32) (sum float64) {
	for _, value := range arr {
		sum += float64(value)
	}
	return sum
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
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
