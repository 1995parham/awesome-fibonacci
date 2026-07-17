# Contributing

Thanks for wanting to add something. This repo collects Fibonacci
implementations across languages and algorithms, so almost any addition is a
good addition — a language that isn't here yet, an algorithm that isn't covered
in a language that is, or tests for an implementation that lacks them.

## The one rule that matters

Every implementation follows the standard indexing:

```
F(0) = 0
F(1) = 1
F(n) = F(n-1) + F(n-2)
```

`F(10)` is `55`. If your implementation says `89`, it's off by one — you've
implemented `F(n+1)`. This is the only thing that has to be consistent across
the whole repo, because comparing implementations is the entire point.

## Layout

- One top-level directory per language, named the way the language is normally
  named: `Haskell`, `JavaScript`, `C++` — not `hs`, `js`, `cpp`.
- Name files after the algorithm they implement: `recursive_fibonacci.py`,
  `matrix.go`, `fast_doubling.rs`.
- Within a directory, use that language's normal project layout and tooling. A
  Rust entry is a Cargo project. A Go entry is a package. A Node entry has a
  `package.json`. Someone who knows the language should find nothing surprising.

## Code

Write it the way the language wants to be written. Idiomatic beats clever, and
clarity beats brevity — these implementations exist to be read and compared, so
a reader who doesn't know your language should still be able to follow the
algorithm.

If the language has a formatter (`gofmt`, `rustfmt`, `black`, `prettier`), run
it. If it has a linter, listen to it.

## Tests

Every language here has them, and CI runs them all. Copy the pattern: a table of
known values checked against every implementation in that language. The Go suite
goes furthest, adding benchmarks.

The table to reuse:

| n | F(n) |
| - | ---- |
| 0 | 0 |
| 1 | 1 |
| 2 | 1 |
| 5 | 5 |
| 10 | 55 |
| 20 | 6765 |
| 30 | 832040 |

Watch out for overflow. `F(93)` no longer fits in a 64-bit unsigned integer, and
`F(47)` overflows a 32-bit one. If your language has bignums, using them is
fine — just say so in a comment.

## Adding a new algorithm

The algorithms table in the [README](README.md#the-algorithms) lists the ones
worth having. If you're implementing something that isn't on it, add a row
explaining the idea and its complexity.

## Submitting

1. Fork and branch.
2. Make your change.
3. Update the implementations table in the [README](README.md#implementations).
4. Open a pull request describing what you added and in what language.

Small PRs are easier to review than big ones. One language or one algorithm per
PR is ideal.
