package main

import (
	"fmt"
)

func main() {
	array2 := [3]int{1, 2, 3}
	array1 := [3]string{"a", "b", "c"}

	var new []interface{}
	for index1, value1 := range array1 {
		new = append(new, value1)
		for index2, value2 := range array2 {
			if index1 == index2 {
				new = append(new, value2)
			}
		}
	}
	fmt.Println(new)

}
