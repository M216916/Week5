package main

import (
	"fmt"
	//"math/big"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func RPN(inputs string) int {
	var stack []int

	arr := strings.Split(inputs, " ")

	for i := range arr {
		if arr[i] == "+" {
			stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
		} else {
			a, _ := strconv.Atoi(arr[i])
			stack = append(stack, a)
		}
	}

	return stack[0]

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputs := scanner.Text()

	output := RPN(inputs)

	fmt.Println(output)
}
