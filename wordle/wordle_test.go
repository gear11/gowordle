package wordle

import (
	"fmt"
	"testing"
)

func Test_eval(t *testing.T) {
	doTestEval(t, "spell", "spill", "ggbgg")
	doTestEval(t, "spoil", "spill", "ggybg")
	doTestEval(t, "rainy", "alien", "ybgby")
	doTestEval(t, "rainy", "saint", "bgggb")
	doTestEval(t, "rainy", "rainy", "ggggg")

	doTestEval(t, "unite", "alien", "bbgyy")
	doTestEval(t, "unite", "brine", "bbgyg")
	doTestEval(t, "unite", "snide", "bggbg")
	doTestEval(t, "unite", "unite", "ggggg")
}

func doTestEval(t *testing.T, actual string, guess string, expected string) {
	colors, _ := eval(actual, guess)
	if colors != expected {
		t.Error("expected ", expected, ", got ", colors)
	}

	filtered := Possible(guess, colors, []string{actual})
	if len(filtered) != 1 {
		t.Error("incorrectly filtered actual ", actual, "from guess ", guess, ", colors", colors)
	}

	fmt.Println(actual, guess, colors, expected, colors == expected, len(filtered) == 1)
}
