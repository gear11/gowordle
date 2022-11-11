# gowordle
My first Go project -- a (non-optimal) solver for Wordle

The intuition for this "brute force" solver is that given a set of possible answer words,
the best next guess is one that eliminates the most possibilities over all of the possible answer words
(since they are presumably equally likely). So, for example, if the possible guesses and answers are `spill`,
`spell`, `sleep`, and `lapse`, we take each of those guess words in turn, play the word against
each answer (so `spill` vs `spill`, `spill` vs `spell`, `spill` vs `sleep`, ...), look at the resulting
colors (`ggggg`, `ggbgg`, `gybyb`, ..), and
then filter the answer list against that guess word / color combination to see how many words
are filtered out. For each guess, sum this over all possible answers. Whichever guess filters out the most possible answers over the answer words
is then recommended.

Over the list of 2315 possible Wordle answers, and a list of 12972 guess words
(fun fact--Wordle allows more guess words than the set of answers, for
example Wordle will never select a plural of a word such as `foxes`, but
will allow it to be played), the solver achieves a guess average of
3.45054 using the recommended starting word `salet`, vs. the proven
limit of ~3.42. Not bad for a very simplistic design!
