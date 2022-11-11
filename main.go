package main

import (
	"fmt"
	"os"
	"wordle/wordle"
)

func main() {
	if len(os.Args) == 1 {
		solveAll()
	} else if len(os.Args)%2 == 0 {
		fmt.Println("Either no arguments or pairs of arguments")
	} else {
		solveNext()
	}

}

func solveNext() {
	fmt.Println("Solving next")
	guessWords, _ := wordle.ReadLines("words2.txt")
	answerWords, _ := wordle.ReadLines("answers.txt")
	var nextGuess string
	for i := 1; i < len(os.Args); i += 2 {
		guess, colors := os.Args[i], os.Args[i+1]
		fmt.Println("Processing", guess, colors)
		answerWords = wordle.Possible(guess, colors, answerWords)
	}
	nextGuess = wordle.BestNextGuess(guessWords, answerWords)
	fmt.Println("Next guess:", nextGuess)
}

func solveAll() {
	fmt.Println("Solving all")
	guessWords, _ := wordle.ReadLines("words2.txt")
	answerWords, _ := wordle.ReadLines("answers.txt")

	fmt.Println("Starting with", len(answerWords), "possibilities")

	sumOfAttempts := 0
	for i, actual := range answerWords {
		guesses := wordle.Solve(actual, guessWords, answerWords, "salet")
		sumOfAttempts += len(guesses)
		avg := float32(sumOfAttempts) / float32(i+1)
		fmt.Println("Solved", actual, "->", guesses, "in", len(guesses), "guesses (", avg, "avg)")

		// With raise + simiple key Solved zonal -> zonal in 5 guesses ( 3.5330453 avg)
		// With hash key Solved zonal -> zonal in 5 guesses ( 3.5287256 avg)
		// With 2 hashes Solved zonal -> zonal in 5 guesses ( 3.5287256 avg)
		// With salet and full word list Solved zonal -> zonal in 3 guesses ( 3.45054 avg)
		// With raise and full word list Solved zonal -> zonal in 4 guesses ( 3.4984882 avg)
		// Solved zonal -> zonal in 3 guesses ( 3.45054 avg)
	}
}
