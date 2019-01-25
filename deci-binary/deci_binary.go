package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func findMaxDeciBin(num int) int {
	if num == 0 {
		return 0
	}

	numList := make([]int, 0)
	for {
		if num == 0 {
			break
		}
		n := num % 10
		num = num / 10
		numList = append(numList, n)
	}

	res := 0
	for i := len(numList) - 1; i >= 0; i-- {
		if numList[i] != 0 {
			res = 10*res + 1
		} else {
			res = 10*res + 0
		}
	}

	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("input error=%v\n", scanner.Text())
			return
		}
		for {
			if input == 0 {
				break
			}
			v := findMaxDeciBin(input)
			fmt.Printf("%v ", v)
			input = input - v
		}
		fmt.Println()
	}
}
