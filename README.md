# gowordle
My first Go project -- a (non-optimal) solver for Wordle

If you run it without any arguments, it will loop through all 2315 possible Wordle answers
using some 12K guess words<sup>*</sup> and show you its solution for each one, keeping a running average:

```bash
Solving all
Starting with 2315 possibilities
Solved aback -> [salet orcin whups aback] in 4 guesses ( 4 avg)
Solved abase -> [salet crash abase] in 3 guesses ( 3.5 avg)
Solved abate -> [salet grate abate] in 3 guesses ( 3.3333333 avg)
Solved abbey -> [salet unrip abbey] in 3 guesses ( 3.25 avg)
Solved abbot -> [salet fiord abbot] in 3 guesses ( 3.2 avg)
```

Otherwise, if you invoke with
pairs of arguments (`guess1`, `colors1`, `guess2`, `colors2`, ...) it will output the next
best guess. For example:

```bash
% ./main alien ybyyy               
Solving next
Next guess: inane
```

The intuition for this "brute force" solver is that given a set of possible answer words,
the best next guess is one that eliminates the most possibilities over all of the possible answer words
(assuming they are equally likely). So, for example, if the possible guesses and answers are `spill`,
`spell`, `sleep`, and `lapse`, we take each of those guess words in turn, play the word against
each answer (so `spill` vs `spill`, `spill` vs `spell`, `spill` vs `sleep`, ...), look at the resulting
colors (`ggggg`, `ggbgg`, `gybyb`, ..), and
then filter the answer list against that guess word / color combination to see how many possible
answers remain (for example, `spill` and `ggbgg` leaves `spell` as the only remaining possibility). 
For each guess / color, sum the lengths of "remain" list over all possible answers. Whichever guess
leaves the fewest possible answers over all possible answer words is then recommended.

Since each next guess tries all possible guesses against all possible answers,
it's an exponential time algorithm (O(n<sup>2</sup>) at least). Raise to O(n<sup>3</sup>)
to solve for all possible answers, and O(n<sup>4</sup>) if we actually wanted to find
the best start word.

Over the list of 2315 possible Wordle answers, and a list of 12972 guess words, the solver achieves a guess average of
3.45054 using the commonly recommended starting word `salet`, vs. the [optimal limit of 3.4201](https://www.poirrier.ca/notes/wordle-optimal/). Not bad for a very simplistic design!

I did add one simple but effective optimization, which is to cache the best next guess for a
given history of guesses and remaining possible answer list. This is most beneficial when
the result is `bbbbb` for the first guess. In that case, the 2nd guess will
always be the same (for a given first guess word), but it takes a while to get there.

A couple of observations about Go in doing this project:

* Go is the first language I've used in a long time that is compiled (to non-bytecode). Some
  time ago I wrote a Wordle solver in Python using a similar algorithm and it is dog slow compared
  to this one. I may try on same hardware with Python 3.11 and see if the gap has closed at all, but 
  it is really nice to have quick compile and fast execution.
* Some things that may not be unique to Go but that I appreciate:
  * Multiple return values
  * The simplicity of `func main() {...` and putting arguments in `os.Args`
  * No parens around the clauses of the `for` loop. I would not have expected this to matter much,
    but the readability is great
  * `for` as a general loop concept vs. `for/while/do` is great. My main gotcha was typing `for x in range` (Python)
    vs. `for x := range` but I take that as just a sign of how natural the range concept is.
  * Opinionated formatting (and Goland's auto reformatting on save)
* Things that grind my gears
  * Go gets knocked for its conventions on error handling and not having exceptions. I will
    pile on here. The lack of `try/catch` leads to a lot of LOC devoted to error
    handling boilerplate. Worse than just boilerplate, if you modify a
    function so that it now has an error condition, you have to modify every call site
    of that function, and potentially cascade the change through the stack.
    Exceptions have drawbacks, but importantly you can add a `throw` (or `raise`) deep
    in the code, and use code reviews and automated testing to assess and verify its
    intended effect without modifying every call site. Go feels like a real downgrade here.
  * The `var` vs. `:=` usage seems quirky. My preference in this area is: always declare first usage
    via `var` or `const`, and only declare types the compiler can't figure out (`Foo x = new Foo()`
    used to bug me a ton in Java)
  * `const` in Go is really weak, and seems more designed for the compiler than the programmer. 
    A compiler-enforced way to declare any value as immutable once assigned/instantiated is
    really valuable for figuring things out. I really like Scala's `var` vs. `val`.
  * Yet another model for package / module semantics. Uppercase as an export signal is
    just goofy, vs. explicit exports. Modifying what a module exports is now a refactor
    vs. editing a declaration.
  * A small thing, but I wish it was `func foo() -> int { ... }` vs. `func foo() int { ... }`. I'm not usually
    for extra typing for things the compiler can figure out, but I find the first easier to
    parse, and `->` is a fairly cross-language convention for indicating the return type of a function.

Overall, I'm glad I did this project to demystify Go. I can understand why it has a following
as a language for microservices, given its fast compile and execution times. 

I can't yet speak to Go as a "programming in the large" language.
Strong typing enforced by the compiler is sure to help; witness how Python and Javascript ->
Typescript evolved type systems to help manage large code bases. Generics seem like a
critical piece as well. While I didn't use generics in this project, I did experiment with them a bit.
In a strongly typed, compiled language like Go, generics are essential to reducing boilerplate
and creating more concise and powerful code. I do think the lack of exceptions is a big deal,
because it creates a lot of error bookkeeping in the code (an interesting attempt to add these
via a [library](https://hackthology.com/exceptions-for-go-as-a-library.html)).

On the topic of error handling, I also like the suggestion [here](https://travix.io/unhandled-errors-in-go-3f341f2704dd)
of a `must` operator or additional keywords that would do sensible things. Maybe on error
I just return the `nil` values for the function and push the `Error` onto a global stack
that someone back in the stack can handle by calling `Sys.except() -> []Error`?

Anyway, I think Wordle itself is a great playground for fun algorithmic puzzles and am
glad that I was able to get a good result out of a pretty dumb algorithm that is
saved by Go's performance.


`*` - fun fact--Wordle allows more guess words than the set of answers, for
example Wordle will never select a plural of a word such as `foxes`, but
will allow it to be played