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
	{0, 1},
	{1, 1},
	{2, 2},
	{3, 3},
	{4, 5},
	{5, 8},
	{6, 13},
	{7, 21},
	{8, 34},
	{9, 55},
	{10, 89},
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
