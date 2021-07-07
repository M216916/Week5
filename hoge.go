package main

import (
	"fmt"
	//"math/big"
	"strconv"
	//"strings"
)

func main() {
	var inputs string

	fmt.Scan(&inputs)

	output, _ := strconv.Atoi(inputs)

	fmt.Println(output)
}
