package main

import "fmt"

func main() {
	var n int
	prevoius := 0
	next := 1
	nextNext := 0
	fmt.Print("Enter the number of turns: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		if i == 1 {
			fmt.Print(" ", prevoius)
			continue
		}
		if i == 2 {
			fmt.Print(" ", next)
			continue
		}
		nextNext = prevoius + next
		prevoius = next
		next = nextNext
		fmt.Print(" ", nextNext)
	}
}
