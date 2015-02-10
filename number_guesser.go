package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Welcome to Number Guesser!")
	number := generateRandomNumber()
	var guess int
	var err error
	for number != guess {
		fmt.Print("Enter your guess: ")
		guess, err = getInputFromUser()
		if err != nil {
			fmt.Println("The guess is not a number.")
		} else {
			evaluateUserGuess(number, guess)
		}
	}
}

func generateRandomNumber() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(100)
}

func getInputFromUser() (int, error) {
	var input string
	fmt.Scanln(&input)
	guess, error := strconv.Atoi(input)
	return guess, error
}

func evaluateUserGuess(number int, guess int) {
	if numberGreaterThanGuess(number, guess) {
		fmt.Println("Your guess is smaller.")
	} else if numberLowerThanGuess(number, guess) {
		fmt.Println("Your guess is greater.")
	} else {
		fmt.Println("Your guess is right!")
	}
}

func numberGreaterThanGuess(number int, guess int) bool {
	return number > guess
}

func numberLowerThanGuess(number int, guess int) bool {
	return number < guess
}
