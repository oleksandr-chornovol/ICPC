package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	inputString, _ := reader.ReadString('\n')
	var input []int
	for _, f := range strings.Fields(inputString) {
		i, _ := strconv.Atoi(f)
		input = append(input, i)
	}

	first := [2]int{input[0], input[1]}
	second := [2]int{input[2], input[3]}
	third := [2]int{input[4], input[5]}

	fToS := math.Abs(float64(first[0]-second[0])) + math.Abs(float64(first[1]-second[1])) - 1
	fToT := math.Abs(float64(first[0]-third[0])) + math.Abs(float64(first[1]-third[1])) - 1
	sToT := math.Abs(float64(second[0]-third[0])) + math.Abs(float64(second[1]-third[1])) - 1

	if fToS < fToT {
		fmt.Println(fToS + min(fToT, sToT))
	} else {
		fmt.Println(fToT + min(fToS, sToT))
	}
}

func min(a float64, b float64) float64 {
	if a < b {
		return a
	}

	return b
}
