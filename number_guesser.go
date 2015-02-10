package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Welcome to Number Guesser!")
	number := generateRandomNumber()
	var guess int
	for number != guess {
		fmt.Print("Enter your guess: ")
		guess = getInputFromUser()
		evaluateUserGuess(number, guess)
	}
}

func generateRandomNumber() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(100)
}

func getInputFromUser() int {
	var guess string
	fmt.Scanln(&guess)
	number, _ := strconv.Atoi(strings.TrimSpace(guess))
	return number
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
