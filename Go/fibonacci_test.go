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

func TestRecursive(t *testing.T) {
	r := New(Recursive)
	for _, test := range tests {
		if got := r.Fibonacci(test.input); got != test.want {
			t.Errorf("Fibonacci(%d) ==> %d != %d)", test.input, got, test.want)
		}
	}
}

func TestLinear(t *testing.T) {
	l := New(Linear)
	for _, test := range tests {
		if got := l.Fibonacci(test.input); got != test.want {
			t.Errorf("Fibonacci(%d) ==> %d != %d)", test.input, got, test.want)
		}
	}
}

func benchmark(t int, i int, b *testing.B) {
	f := New(t)
	for n := 0; n < b.N; n++ {
		f.Fibonacci(i)
	}
}

func BenchmarkRecursive10(b *testing.B) { benchmark(Recursive, 10, b) }
func BenchmarkLinear10(b *testing.B)    { benchmark(Linear, 10, b) }
