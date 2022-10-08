package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	reader := bufio.NewReader(os.Stdin)
	inputString, _ := reader.ReadString('\n')
	var input []int
	for _, f := range strings.Fields(inputString) {
		i, _ := strconv.Atoi(f)
		input = append(input, i)
	}

	arrowsCount := 0
	for ; n > 0; arrowsCount++ {
		height := input[0]

		for i := 0; i < len(input); i++ {
			if height == input[i] {
				input = append(input[:i], input[i+1:]...)
				n--
				i--
				height--
			}
			if height == 0 || n == 0 {
				break
			}
		}
	}

	fmt.Println(arrowsCount)
}
