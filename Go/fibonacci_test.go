/*
 * In The Name Of God
 * ========================================
 * [] File Name : fibonacci_test.go
 *
 * [] Creation Date : 22-11-2015
 *
 * [] Created By : Parham Alvani (parham.alvani@gmail.com)
 * =======================================
 */
/*
 * Copyright (c) 2015 Parham Alvani.
 */

package fibonacci

import (
	"testing"
)

var tests = []struct {
	input int
	want  int
}{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
	{8, 21},
	{9, 34},
	{10, 55},
	{20, 6765},
	{30, 832040},
}

var implementations = []struct {
	name string
	kind int
}{
	{"Recursive", Recursive},
	{"Memoized", Memoized},
	{"Linear", Linear},
	{"Matrix", Matrix},
	{"FastDoubling", FastDoubling},
	{"ClosedForm", ClosedForm},
}

// TestImplementations runs every implementation against the same table. They
// only earn the right to be compared if they agree.
func TestImplementations(t *testing.T) {
	for _, impl := range implementations {
		t.Run(impl.name, func(t *testing.T) {
			f := New(impl.kind)
			for _, test := range tests {
				if got := f.Fibonacci(test.input); got != test.want {
					t.Errorf("Fibonacci(%d) ==> %d != %d)", test.input, got, test.want)
				}
			}
		})
	}
}

// TestAgreement checks the implementations against each other rather than
// against the table, over a contiguous range the table only samples.
func TestAgreement(t *testing.T) {
	// Bounded by ClosedForm, which is exact only up to closedFormExactUpTo,
	// and by Recursive, which becomes unbearably slow well before that.
	for n := range 30 {
		want := New(Linear).Fibonacci(n)
		for _, impl := range implementations {
			if got := New(impl.kind).Fibonacci(n); got != want {
				t.Errorf("%s.Fibonacci(%d) ==> %d != %d", impl.name, n, got, want)
			}
		}
	}
}

// closedFormExactUpTo is the largest n for which Binet's formula still rounds
// to the right integer in float64.
const closedFormExactUpTo = 75

// TestClosedFormPrecisionLimit pins down where the closed form stops being
// exact. This is a documented property rather than a bug, but a silent change
// to it should still fail the build.
func TestClosedFormPrecisionLimit(t *testing.T) {
	c, l := New(ClosedForm), New(Linear)

	for n := range closedFormExactUpTo + 1 {
		if got, want := c.Fibonacci(n), l.Fibonacci(n); got != want {
			t.Errorf("closed form should be exact at n=%d: %d != %d", n, got, want)
		}
	}

	n := closedFormExactUpTo + 1
	if c.Fibonacci(n) == l.Fibonacci(n) {
		t.Errorf("closed form is now exact at n=%d; the documented limit moved", n)
	}
}

// benchmark builds the implementation inside the loop on purpose. Memoized
// keeps its cache between calls, so hoisting the New out would measure a map
// lookup rather than the algorithm.
func benchmark(t int, i int, b *testing.B) {
	for range b.N {
		New(t).Fibonacci(i)
	}
}

func BenchmarkRecursive10(b *testing.B)    { benchmark(Recursive, 10, b) }
func BenchmarkMemoized10(b *testing.B)     { benchmark(Memoized, 10, b) }
func BenchmarkLinear10(b *testing.B)       { benchmark(Linear, 10, b) }
func BenchmarkMatrix10(b *testing.B)       { benchmark(Matrix, 10, b) }
func BenchmarkFastDoubling10(b *testing.B) { benchmark(FastDoubling, 10, b) }
func BenchmarkClosedForm10(b *testing.B)   { benchmark(ClosedForm, 10, b) }

func BenchmarkRecursive30(b *testing.B)    { benchmark(Recursive, 30, b) }
func BenchmarkMemoized30(b *testing.B)     { benchmark(Memoized, 30, b) }
func BenchmarkLinear30(b *testing.B)       { benchmark(Linear, 30, b) }
func BenchmarkMatrix30(b *testing.B)       { benchmark(Matrix, 30, b) }
func BenchmarkFastDoubling30(b *testing.B) { benchmark(FastDoubling, 30, b) }
func BenchmarkClosedForm30(b *testing.B)   { benchmark(ClosedForm, 30, b) }
