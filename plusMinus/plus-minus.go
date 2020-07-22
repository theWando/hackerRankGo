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
Complete the plusMinus function below.
https://www.hackerrank.com/challenges/plus-minus/problem?h_r=next-challenge&h_v=zen
*/

func plusMinus(arr []int32) {
	p, n, z := plusMinusWork(arr)
	fmt.Printf("%.6f\n", p)
	fmt.Printf("%.6f\n", n)
	fmt.Printf("%.6f\n", z)

}

func plusMinusWork(arr []int32) (float64, float64, float64) {
	var (
		p, n, z float64
	)

	for _, val := range arr {
		if val == 0 {
			z++
		} else if val < 0 {
			n++
		} else if val > 0 {
			p++
		}
	}
	length := float64(len(arr))
	return math.Round(p/length*1000000) / 1000000, math.Round(n/length*1000000) / 1000000, math.Round(z/length*1000000) / 1000000
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
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
