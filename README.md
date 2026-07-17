# Awesome Fibonacci

[![CI](https://github.com/1995parham/awesome-fibonacci/actions/workflows/ci.yml/badge.svg)](https://github.com/1995parham/awesome-fibonacci/actions/workflows/ci.yml)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](CONTRIBUTING.md)

> The Fibonacci sequence, in every language, by every algorithm.

A collection of Fibonacci implementations written to be *read and compared*, not
just to run. The same problem, solved in different languages and with different
algorithms, is a surprisingly good lens on what each language and each approach
actually costs you — in speed, in memory, and in how much code you have to write.

## Contents

- [The Sequence](#the-sequence)
- [Implementations](#implementations)
- [The Algorithms](#the-algorithms)
- [Running Them](#running-them)
- [Contributing](#contributing)
- [License](#license)

## The Sequence

Every implementation here should follow the standard indexing:

```
F(0) = 0
F(1) = 1
F(n) = F(n-1) + F(n-2)   for n > 1
```

Giving: `0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...`

Every implementation is checked against the same table of known values in CI, so
`F(10)` is `55` in all of them. That consistency is what makes comparing them
worth anything.

## Implementations

| Language | Recursive | Linear | Logarithmic | Tests |
| -------- | :-------: | :----: | :---------: | :---: |
| [Go](Go/) | ✅ | ✅ | ❌ | ✅ |
| [Lisp](Lisp/) | ✅ | ❌ | ❌ | ✅ |
| [Python](Python/) | ✅ | ❌ | ❌ | ✅ |
| [Rust](Rust/) | ✅ | ❌ | ❌ | ✅ |

Languages we'd love to see: C, C++, Haskell, JavaScript, Ruby, Elixir, Zig,
Clojure, OCaml, Scala, Kotlin, Swift, Erlang, Prolog, APL — and anything else
you can make a compiler accept.

## The Algorithms

Fibonacci is a small problem with a genuinely wide spread of solutions, which is
what makes it worth collecting:

| Algorithm | Time | Space | Idea |
| --------- | ---- | ----- | ---- |
| **Recursive** | `O(φⁿ)` | `O(n)` | Transcribe the definition directly. Recomputes the same subproblems exponentially many times. |
| **Memoized** | `O(n)` | `O(n)` | The recursive version, but cache each `F(n)` the first time you compute it. |
| **Linear** | `O(n)` | `O(1)` | Iterate upward keeping only the last two values. |
| **Matrix power** | `O(log n)` | `O(1)` | `[[1,1],[1,0]]ⁿ` has `F(n)` in it; exponentiate by squaring. |
| **Fast doubling** | `O(log n)` | `O(1)` | `F(2k)` and `F(2k+1)` derive from `F(k)` and `F(k+1)` directly. Matrix power without the matrix. |
| **Closed form** | `O(1)`\* | `O(1)` | Binet's formula: `F(n) = (φⁿ - ψⁿ) / √5`. Loses precision to floating point surprisingly early. |

\* Assuming constant-time `pow`, which stops being true once the numbers outgrow
a machine word — as they do around `F(93)` for 64-bit integers.

## Running Them

Each language lives in its own top-level directory and is self-contained.

```sh
# Go — tests plus benchmarks
cd Go && go test -bench=.

# Rust — `cargo run` prompts for a length and prints the sequence
cd Rust && cargo test

# Python
cd Python && python -m unittest discover

# Lisp
cd Lisp && sbcl --script fibonacci_test.lisp
```

## Contributing

Contributions are very welcome, whether that's a new language, a new algorithm
in a language that's already here, or tests for one that lacks them.

The short version:

1. Put your code in the directory for its language, creating it if it doesn't
   exist yet (use the language's conventional name — `Haskell`, not `hs`).
2. Follow the standard indexing: `F(0) = 0`, `F(1) = 1`.
3. Name files after the algorithm — `recursive_fibonacci.py`, `matrix.go`.
4. Use the language's idiomatic layout and tooling. A Rust entry should be a
   Cargo project; a Go entry should be a package with `_test.go` files.
5. Add tests if the language makes it reasonable to.
6. Update the table in [Implementations](#implementations).

See [CONTRIBUTING.md](CONTRIBUTING.md) for the longer version.

## License

[GPL-3.0](LICENSE) © Parham Alvani
