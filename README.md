# gowordle
My first Go project -- a (non-optimal) solver for Wordle

The intuition for this "brute force" solver is that given a set of possible answer words,
the best next guess is one that eliminates the most possibilities over all of the possible answer words
(assuming they are equally likely). So, for example, if the possible guesses and answers are `spill`,
`spell`, `sleep`, and `lapse`, we take each of those guess words in turn, play the word against
each answer (so `spill` vs `spill`, `spill` vs `spell`, `spill` vs `sleep`, ...), look at the resulting
colors (`ggggg`, `ggbgg`, `gybyb`, ..), and
then filter the answer list against that guess word / color combination to see how many words
are filtered out. For each guess / color, sum the "residuals" over all possible answers. Whichever guess
filters out the most possible answers over the possible answer words is then recommended.

Over the list of 2315 possible Wordle answers, and a list of 12972 guess words*, the solver achieves a guess average of
3.45054 using the recommended starting word `salet`, vs. the proven
limit of ~3.42. Not bad for a very simplistic design!

I did add one simple but effective optimization, which is to cache the best next guess for a
given history of guesses and remaining possible answer list. This is especially impactful because
there will always be a lot of answer words that return `bbbbb` for the first guess (regardless of your
first guess word). In that case, your 2nd guess is always the same, right?

A couple of observations in doing this project:

* Go is the first language I've used in a long time that is compiled (other than a JVM language). Some
  time ago I wrote a Wordle solver in Python using a similar algorithm and it is dog slow compared
  to this one. I may try on same hardware with Python 3.11 and see if the gap has closed at all, but 
  it is really nice to have quick compile and fast execution.
* Go gets knocked for its conventions on error handling and not having exceptions. I will pile
  on here. Yes exceptions interrupt the flow of control, but just being able to bubble up to
  the best layer to handle the issue--so much better than this. `Throw` / `catch` is not ideal, but
  like democracy its the best thing we've got. Go feels like a real downgrade here.
* Some things that may not be unique to Go but that I appreciate:
  * Multiple return values
  * No parens around the clauses of the `for` loop. I would not have expected this to matter much,
    but the readability is great
  * `for` as a general loop concept vs. `for/while/do` is great. My main gotcha was typing `for x in range` (Python)
    vs. `for x := range` but I take that as just a sign of how natural the range concept is.
  * Opinionated formatting (and Goland's auto reformatting on save)
* Grinds my gears
  * The `var` vs. `:=` usage seems quirky. My preference in this area is: always declare first usage
    via `var` or `const`, and only declare types the compiler can't figure out (`Foo x = new Foo()`
    used to bug me a ton in Java)
  * `const` in Go is really weak, and seems more designed for the compiler than the programmer. 
    A compiler-enforced way to declare any value as immutable once assigned/instantiated is
    really valuable for figuring things out. I really like Scala's `var` vs. `val`.
  * Yet another model for package / module semantics. Uppercase as an export signal is
    just goofy, vs. explicit exports.
  * The lack of `try/catch` leads to a lot of LOC devoted to boilerplate types of error
    handling. I love the suggestion [here](https://travix.io/unhandled-errors-in-go-3f341f2704dd)
    of a `must` operator or additional keywords that would do sensible things. Maybe on error
    I just return the `nil` values for the function and push the `Error` onto a global stack
    that someone back in the stack can handle by calling `Sys.except() -> []Error`?
  * I wish it was `func foo() -> int { ... }` vs. `func foo() int { ... }`. I never would have
    imagined pining for superfluous characters in a language for common declarations, but I feel like `->` has
    become a great cross-language convention for indicating return type of a function.

Overall, I'm glad I did this project to demystify Go. I can understand why it has a following
as a language for microservices, given its fast compile and execution times. 

I am inconclusive on Go as a "programming in the large" language.
On the one had, Go is a "small" language, and "small" languages like Python and Javascript evolved to programming in the large
by adding optional typing. But Go already has that, does it need anything? Probably more flexibility.
While I didn't use generics in this project, I did experiment with them a bit, and can see
how in a strongly typed, compiled language like Go, generics are essential to reducing boilerplate
and creating more powerful code.

`*` - fun fact--Wordle allows more guess words than the set of answers, for
example Wordle will never select a plural of a word such as `foxes`, but
will allow it to be played