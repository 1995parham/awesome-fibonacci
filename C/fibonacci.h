#ifndef FIBONACCI_H
#define FIBONACCI_H

#define ERRNO -1

typedef long long int64;

enum fib_algorithm {
  RECURSIVE,
  MEMOIZED,
  LINEAR,
  MATRIX,
  FAST_DOUBLING,
  CLOSED_FORM,
};

int64 fibonacci(enum fib_algorithm algorithm, int n);

#endif
