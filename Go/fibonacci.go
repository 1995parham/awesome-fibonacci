/*
 * In The Name Of God
 * ========================================
 * [] File Name : fibonacci.go
 *
 * [] Creation Date : 28-05-2015
 *
 * [] Created By : Parham Alvani (parham.alvani@gmail.com)
 * =======================================
 */
/*
 * Copyright (c) 2015 Parham Alvani.
 */

package fibonacci

// Fibonacci implementation constants
const (
	Recursive = iota
	Linear
	Logarithmic
)

// Implementation provides abstraction for fibonacci implmeentations
type Implementation interface {
	Fibonacci(int) int
}

// New returns asked implementation of fibonacci
func New(t int) Implementation {
	switch t {
	case Recursive:
		return new(recursive)
	case Linear:
		return new(linear)
	default:
		return nil
	}
}

type recursive struct{}

func (r *recursive) Fibonacci(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return r.Fibonacci(n-1) + r.Fibonacci(n-2)
}

type linear struct{}

func (r *linear) Fibonacci(n int) int {
	a := 1
	b := 1

	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			a += b
		} else {
			b += a
		}
	}

	if n%2 == 0 {
		return a
	}

	return b
}
