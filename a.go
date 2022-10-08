package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	reader := bufio.NewReader(os.Stdin)

	var input []string
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')

		input = append(input, line)
	}

	for _, line := range input {
		checkPair(line)
	}
}

var keyboard = [3][10]string{
	{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p"},
	{"a", "s", "d", "f", "g", "h", "j", "k", "l"},
	{"z", "x", "c", "v", "b", "n", "m"},
}

func checkPair(line string) {
	if line[0] == line[2] {
		fmt.Println("Yes")
		return
	}

	for _, row := range keyboard {
		isHere := false
		for _, key := range row {
			if key == string(line[0]) || key == string(line[2]) {
				if isHere {
					fmt.Println("Yes")
					return
				} else {
					isHere = true
				}
			}
		}
	}

	fmt.Println("No")
}
