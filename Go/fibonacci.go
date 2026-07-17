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

import "math"

// Fibonacci implementation constants
const (
	Recursive = iota
	Memoized
	Linear
	Matrix
	FastDoubling
	ClosedForm
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
	case Memoized:
		return newMemoized()
	case Linear:
		return new(linear)
	case Matrix:
		return new(matrix)
	case FastDoubling:
		return new(fastDoubling)
	case ClosedForm:
		return new(closedForm)
	default:
		return nil
	}
}

type recursive struct{}

func (r *recursive) Fibonacci(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return r.Fibonacci(n-1) + r.Fibonacci(n-2)
}

// memoized is the recursive implementation with each F(n) cached the first
// time it is computed, which collapses the exponential call tree to O(n).
// It carries state, so unlike the others an instance is worth reusing.
type memoized struct {
	cache map[int]int
}

func newMemoized() *memoized {
	return &memoized{cache: map[int]int{0: 0, 1: 1}}
}

func (m *memoized) Fibonacci(n int) int {
	if f, ok := m.cache[n]; ok {
		return f
	}

	f := m.Fibonacci(n-1) + m.Fibonacci(n-2)
	m.cache[n] = f

	return f
}

type linear struct{}

func (r *linear) Fibonacci(n int) int {
	a := 0
	b := 1

	for range n {
		a, b = b, a+b
	}

	return a
}

// mat2 is a 2x2 matrix laid out in row-major order.
type mat2 [2][2]int

func (m mat2) mul(o mat2) mat2 {
	return mat2{
		{m[0][0]*o[0][0] + m[0][1]*o[1][0], m[0][0]*o[0][1] + m[0][1]*o[1][1]},
		{m[1][0]*o[0][0] + m[1][1]*o[1][0], m[1][0]*o[0][1] + m[1][1]*o[1][1]},
	}
}

// matrix raises [[1,1],[1,0]] to the n'th power, which puts F(n) in the
// top-right cell, using exponentiation by squaring to do it in O(log n).
type matrix struct{}

func (m *matrix) Fibonacci(n int) int {
	result := mat2{{1, 0}, {0, 1}}
	base := mat2{{1, 1}, {1, 0}}

	for e := n; e > 0; e >>= 1 {
		if e&1 == 1 {
			result = result.mul(base)
		}

		base = base.mul(base)
	}

	return result[0][1]
}

// fastDoubling derives F(2k) and F(2k+1) from F(k) and F(k+1) directly:
//
//	F(2k)   = F(k) * (2*F(k+1) - F(k))
//	F(2k+1) = F(k)^2 + F(k+1)^2
//
// This is the matrix method with the matrix algebra folded out by hand, so it
// is also O(log n) but without the redundant multiplications.
type fastDoubling struct{}

func (f *fastDoubling) Fibonacci(n int) int {
	a, _ := f.pair(n)

	return a
}

// pair returns F(n) and F(n+1) together, halving n on each step.
func (f *fastDoubling) pair(n int) (int, int) {
	if n == 0 {
		return 0, 1
	}

	a, b := f.pair(n / 2)
	c := a * (2*b - a)
	d := a*a + b*b

	if n%2 == 0 {
		return c, d
	}

	return d, c + d
}

// closedForm evaluates Binet's formula. It is the only implementation here
// that does not compute F(n) exactly: it rounds a float64 to the nearest
// integer, and the rounding error in math.Pow compounds with n until it
// exceeds 1/2 and the rounding lands on the wrong integer. That happens at
// n=76, well before F(n) itself outgrows a float64's 53-bit mantissa, so the
// limit here is precision and not range. See TestClosedFormPrecisionLimit.
type closedForm struct{}

func (c *closedForm) Fibonacci(n int) int {
	sqrt5 := math.Sqrt(5)
	phi := (1 + sqrt5) / 2
	psi := (1 - sqrt5) / 2

	return int(math.Round((math.Pow(phi, float64(n)) - math.Pow(psi, float64(n))) / sqrt5))
}
