package main

import (
	"fmt"
	"strings"
)

func main() {
	diamond(10)
}
func diamond(x int) {
	for i := 0; i < x; i++ {
		fmt.Println(strings.Repeat(" ", x-i-1) + strings.Repeat("x ", i+1))
	}
	for j := x - 1; j > 0; j-- {
		fmt.Println(strings.Repeat(" ", x-j) + strings.Repeat("x ", j))
	}
}
