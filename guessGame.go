package main

import (
	"fmt"
	"math/rand"
)

func main() {
	random := rand.Intn(20)
	for true {
		var guess int
		fmt.Println("Enter your guess...")
		_, err := fmt.Scan(&guess)

		if err != nil {
			fmt.Println("Something Went Wrong....")
		}

		fmt.Println("random", random)
		guesses := []int{}

		if guess > random {
			if !contains(guesses, guess) {
				guesses = append(guesses, guess)
			}
			fmt.Println("Too hign")
			fmt.Println("You have tried", len(guesses), "times")
		}
		if guess < random {
			if !contains(guesses, guess) {
				guesses = append(guesses, guess)
			}
			fmt.Println("Too low")
			fmt.Println("You have tried", len(guesses), "times")
		}
		if guess == random {
			fmt.Println("Congratulation")
			break
		}
	}

}
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
