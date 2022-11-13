package main

import (
	"bufio"
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

func MustReadLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = file.Close()
	}()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func solveNext() {
	fmt.Println("Solving next")
	guessWords := MustReadLines("data/words2.txt")
	answerWords := MustReadLines("data/answers.txt")
	var nextGuess string
	for i := 1; i < len(os.Args); i += 2 {
		guess, colors := os.Args[i], os.Args[i+1]
		answerWords = wordle.Possible(guess, colors, answerWords)
	}
	nextGuess = wordle.BestNextGuess(guessWords, answerWords)
	fmt.Println("Next guess:", nextGuess)
}

func solveAll() {
	fmt.Println("Solving all")
	guessWords := MustReadLines("data/words2.txt")
	answerWords := MustReadLines("data/answers.txt")

	fmt.Println("Starting with", len(answerWords), "possibilities")

	sumOfAttempts := 0
	for i, actual := range answerWords {
		guesses := wordle.Solve(actual, guessWords, answerWords, "salet")
		sumOfAttempts += len(guesses)
		avg := float32(sumOfAttempts) / float32(i+1)
		fmt.Println("Solved", actual, "->", guesses, "in", len(guesses), "guesses (", avg, "avg)")
	}
}
