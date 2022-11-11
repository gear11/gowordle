package wordle

import (
	"fmt"
	"hash/fnv"
	"math"
)

// eval evaluates the given guess against the actual string
// and returns a color string of the form 'ggbyb' or whatver
func eval(actual string, guess string) (string, error) {
	gl := len(guess)
	r := make([]rune, gl)
	if gl != len(actual) {
		return "", fmt.Errorf("guess %s and actual %s of different lengths", guess, actual)
	}
	gu, ac := []rune(guess), []rune(actual)
	// Pass 1 for green
	for i := 0; i < gl; i++ {
		if gu[i] == ac[i] {
			r[i] = 'g'
			ac[i] = '_'
		}
	}
	// Pass 2 for yellow or black
	for i := 0; i < gl; i++ {
		if r[i] == 'g' {
			continue
		}
		found := false
		for j := 0; j < gl; j++ {
			if gu[i] == ac[j] {
				r[i] = 'y'
				ac[j] = '_'
				found = true
			}
		}
		if !found {
			r[i] = 'b'
		}
	}
	return string(r), nil
}

// isPossible indicates for the given played string and resulting colors,
// whether the given word is still a possibility
func isPossible(played string, colors string, word string) bool {
	pl := len(played)
	pr, cr, wr := []rune(played), []rune(colors), []rune(word)

	// Need to do in 3 passes with replacement, to handle played with repeat letters
	// Green
	for i := 0; i < pl; i++ {
		p, c, w := pr[i], cr[i], wr[i]
		if c == 'g' {
			if p != w {
				return false
			}
			wr[i] = '_' // Prevent additional matches
		}
	}
	// Yellow
	for i := 0; i < pl; i++ {
		p, c, w := pr[i], cr[i], wr[i]
		if c == 'y' {
			if p == w {
				return false // Would have been green if matched
			}
			found := false
			for j := 0; j < pl; j++ {
				if p == wr[j] {
					wr[j] = '_' // Prevent additional matches
					found = true
				}
			}
			if !found {
				return false
			}
		}
	}
	// Black
	for i := 0; i < pl; i++ {
		p, c := pr[i], cr[i]
		if c == 'b' {
			for j := 0; j < pl; j++ {
				if p == wr[j] {
					return false
				}
			}
		}
	}
	return true
}

// Possible filters the given slice of words based on the given guess and resulting color
// and returns which words in the given slice are still possible
func Possible(guess string, colors string, words []string) []string {
	var all []string
	for _, word := range words {
		if isPossible(guess, colors, word) {
			all = append(all, word)
		}
	}
	return all
}

func BestNextGuess(guessWords, answerWords []string) string {
	var bestGuess string
	var leastRemaining = math.MaxInt32
	var guessInAnswers bool // Prefer guesses that are among Possible answers

	for _, guess := range guessWords {
		remaining := 0
		inAnswers := false
		for _, actual := range answerWords {
			if guess == actual {
				remaining += 1
				inAnswers = true
			} else {
				colors, _ := eval(actual, guess)
				remaining += len(Possible(guess, colors, answerWords))
			}
		}
		if (remaining < leastRemaining) || (remaining == leastRemaining && !guessInAnswers) {
			bestGuess = guess
			leastRemaining = remaining
			guessInAnswers = inAnswers
		}
	}
	return bestGuess
}

type key struct {
	guesses   uint32
	remaining uint32
}

var cachedGuesses = make(map[key]string)

func hash(s string) uint32 {
	h := fnv.New32a()
	_, _ = h.Write([]byte(s))
	return h.Sum32()
}

func hashSlice(ss []string) uint32 {
	var h uint32 = 0
	for _, s := range ss {
		h = h ^ hash(s)
	}
	return h
}

func cachedGuess(guesses, remaining []string) (cached string, ok bool) {
	hg, hr := hashSlice(guesses), hashSlice(remaining)
	cached, ok = cachedGuesses[key{hg, hr}]
	return cached, ok
}

func cacheGuess(guess string, guesses, remaining []string) {
	hg, hr := hashSlice(guesses), hashSlice(remaining)
	cachedGuesses[key{hg, hr}] = guess
}

func Solve(actual string, guessWords, answerWords []string, guess string) (guesses []string) {

	remaining := answerWords
	colors := ""

	for colors != "ggggg" {
		if len(guesses) != 0 {
			cached, ok := cachedGuess(guesses, remaining)
			if ok {
				guess = cached
			} else {
				guess = BestNextGuess(guessWords, remaining)
				cacheGuess(guess, guesses, remaining)
			}
		}
		guesses = append(guesses, guess)
		colors, _ = eval(actual, guess)
		remaining = Possible(guess, colors, remaining)
	}
	return guesses
}
