package main

import (
	"fmt"
	//"math/big"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func delete(slice []int, i int) (int, []int) {
	ret := slice[i]
	if len(slice) < i || len(slice) < i {
		return 0, nil
	}
	ans := make([]int, len(slice))
	copy(ans, slice)

	ans = append(slice[:i], slice[(i+1):]...)

	return ret, ans
}

func RPN(inputs string) int {
	var stack []int
	var x, y int

	arr := strings.Split(inputs, " ")

	for i := range arr {
		if arr[i] == "+" {
			y, stack = delete(stack, len(stack)-1)
			x, stack = delete(stack, len(stack)-1)
			stack = append(stack, x+y)
		} else if arr[i] == "-" {
			y, stack = delete(stack, len(stack)-1)
			x, stack = delete(stack, len(stack)-1)
			stack = append(stack, x-y)
		} else if arr[i] == "*" {
			y, stack = delete(stack, len(stack)-1)
			x, stack = delete(stack, len(stack)-1)
			stack = append(stack, x*y)
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
