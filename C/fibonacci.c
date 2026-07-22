/**
 * fibonacci.c
 *
 * The Fibonacci sequence by all 6 algorithms in C.
 */

#include "fibonacci.h"
#include "cc.h" // https://github.com/JacksonAllan/CC

#include <math.h>
#include <string.h>

/**
 * recursive
 */

// fibonacci_recursive is a recursive version of Fibonacci computation.
static int64 fibonacci_recursive(int n) {
  if (n < 0)
    return ERRNO;
  if (n == 0)
    return 0;
  if (n == 1)
    return 1;

  return fibonacci_recursive(n - 1) + fibonacci_recursive(n - 2);
}

/**
 * memoized
 */

// fibonacci_memoized is an improvement to fibonacci_recursive,
// using a hash map to cache computed Fibonacci values to avoid
// duplicated computing.
static int64 fibonacci_memoized_worker(map(int, int64) * cache, int n) {
  if (n < 2)
    return n;

  int64 *value = get(cache, n);
  if (value != NULL) {
    return *value;
  }

  int64 first = fibonacci_memoized_worker(cache, n - 1);
  int64 second = fibonacci_memoized_worker(cache, n - 2);
  if (first == ERRNO || second == ERRNO)
    return ERRNO;

  int64 ret = first + second;

  // cache computed value
  if (!insert(cache, n, ret)) {
    fprintf(stderr, "failed to insert into map\n");
    return ERRNO;
  }

  return ret;
}

static int64 fibonacci_memoized(int n) {
  if (n < 0)
    return ERRNO;

  // Using github.com/JacksonAllan/CC, where CC stands for Convenient
  // Containers. This single-header project provisions a set of modern hash map
  // API.
  map(int, int64) cache;
  init(&cache);

  int64 ret = fibonacci_memoized_worker(&cache, n);
  cleanup(&cache);
  return ret;
}

/**
 * linear
 */

// fibonacci_linear is an iteration version of Fibonacci computing algorithms.
// This version leverages two variables to store values needed for computation
// in a dynamic programming way, thus reducing space complexity to O(1).
static int64 fibonacci_linear(int n) {
  int64 last = 0, curr = 1;
  for (int i = 0; i < n; i++) {
    int64 temp = curr;
    curr += last;
    last = temp;
  }

  return last;
}

/**
 * matrix
 */

typedef struct {
  int64 data[2][2];
} mat2_t;

static mat2_t mul(mat2_t mat_A, mat2_t mat_B) {
  return (mat2_t){
      .data =
          {
              {mat_A.data[0][0] * mat_B.data[0][0] +
                   mat_A.data[0][1] * mat_B.data[1][0],
               mat_A.data[0][0] * mat_B.data[0][1] +
                   mat_A.data[0][1] * mat_B.data[1][1]},
              {mat_A.data[1][0] * mat_B.data[0][0] +
                   mat_A.data[1][1] * mat_B.data[1][0],
               mat_A.data[1][0] * mat_B.data[0][1] +
                   mat_A.data[1][1] * mat_B.data[1][1]},
          },
  };
}

// fibonacci_matrix raises the Fibonacci transformation matrix {{1, 1}, {1, 0}}
// to the n-th power using exponentiation by squaring. F(n) is stored in the
// top-right cell of the resulting matrix, giving O(log n) time complexity.
static int64 fibonacci_matrix(int n) {
  mat2_t result = {
      .data = {{1, 0}, {0, 1}},
  };
  mat2_t base = {
      .data = {{1, 1}, {1, 0}},
  };

  for (int e = n; e > 0; e >>= 1) {
    if ((e & 1) == 1) {
      result = mul(result, base);
    }
    base = mul(base, base);
  }
  return result.data[0][1];
}

/**
 * fast doubling
 */

typedef struct {
  int64 first, second;
} tuple;

// pair returns F(n) and F(n+1) together. It recursively computes the pair
// for n/2, then derives the requested pair with the fast-doubling identities.
static tuple pair(int n) {
  if (n == 0)
    return (tuple){0, 1};

  tuple half = pair(n / 2);
  int64 a = half.first, b = half.second;

  int64 c = a * (2 * b - a);
  int64 d = a * a + b * b;

  if (n % 2 == 0)
    return (tuple){c, d};

  return (tuple){d, c + d};
}

// fibonacci_fast_doubling computes Fibonacci numbers from the identities
// F(2k) = F(k) * (2 * F(k + 1) - F(k)) and
// F(2k + 1) = F(k)^2 + F(k + 1)^2, giving O(log n) time complexity.
static int64 fibonacci_fast_doubling(int n) { return pair(n).first; }

/**
 * closed form
 */

// fibonacci_closed_form evaluates Binet's formula using the golden ratio and
// its conjugate. The result is rounded to the nearest integer, but floating-
// point precision means this implementation is not exact for sufficiently
// large values of n.
static int64 fibonacci_closed_form(int n) {
  double phi = (1 + sqrt(5)) / 2;
  double psi = (1 - sqrt(5)) / 2;

  return (int64)(round((pow(phi, (double)n) - pow(psi, (double)n)) / sqrt(5)));
}

// fibonacci is the main entrance for all implementations.
int64 fibonacci(enum fib_algorithm algorithm, int n) {
  switch (algorithm) {
  case RECURSIVE:
    return fibonacci_recursive(n);
  case MEMOIZED:
    return fibonacci_memoized(n);
  case LINEAR:
    return fibonacci_linear(n);
  case MATRIX:
    return fibonacci_matrix(n);
  case FAST_DOUBLING:
    return fibonacci_fast_doubling(n);
  case CLOSED_FORM:
    return fibonacci_closed_form(n);
  default:
    return ERRNO;
  }
}
