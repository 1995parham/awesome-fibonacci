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
	{2, 2},
	{3, 3},
	{4, 5},
	{5, 8},
	{6, 13},
}

func TestRecursive(t *testing.T) {
	r := New(Recursive)
	for _, test := range tests {
		if got := r.Fibonacci(test.input); got != test.want {
			t.Errorf("Fibonacci(%d) ==> %d != %d)", test.input, got, test.want)
		}
	}
}
