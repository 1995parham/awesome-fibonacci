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

// Implementation provides abstraction for fibonacci implmeentations
type Implementation interface {
	Fibonacci(int)
}

// New returns
func New() Implementation {
}

type recursive struct{}

func (r *recursive) Fibonacci(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return r.Fibonacci(n-1) + r.Fibonacci(n-2)
}
