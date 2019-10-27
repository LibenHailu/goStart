package main

import (
	"fmt"
)

func main() {
	for j := 1; j <= 12; j++ {
		fmt.Println("  ", j, "*")
		for i := 1; i <= 10; i++ {
			fmt.Println(j, "*", i, "=", i*j)

		}
	}
}
