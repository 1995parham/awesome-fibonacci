/**
 * fibonacci_test.c
 *
 * Test file for fibonacci.c.
 * To run the tests:
 *
 * ```bash
 *     # Run the following commands at the root directory of this project
 *     cc -std=c11 C/fibonacci.c C/fibonacci_test.c -lm -o /tmp/fibonacci_test
 *     /tmp/fibonacci_test
 * ```
 *
 */
#include <stdio.h>
#include <stdlib.h>

#include "fibonacci.h"

struct test_case {
  int input;
  int64 expected;
};

struct implementation {
  const char *name;
  int mask;
};

static const struct test_case test_cases[] = {
    {0, 0}, {1, 1}, {2, 1}, {5, 5}, {10, 55}, {20, 6765}, {30, 832040},
};

static const struct implementation implementations[] = {
    {"recursive", RECURSIVE},
    {"memoized", MEMOIZED},
    {"linear", LINEAR},
    {"matrix", MATRIX},
    {"fast doubling", FAST_DOUBLING},
    {"closed form", CLOSED_FORM},
};

static const int closed_form_exact_up_to = 70;

static int test_closed_form_precision_limit(void) {
  int failures = 0;
  for (int n = 0; n <= closed_form_exact_up_to; ++n) {
    int64 got = fibonacci(CLOSED_FORM, n);
    int64 expected = fibonacci(LINEAR, n);

    if (got != expected) {
      fprintf(stderr, "closed form should be exact at n=%d: %lld != %lld", n,
              got, expected);
      ++failures;
    }
  }

  int n = closed_form_exact_up_to + 1;
  if (fibonacci(CLOSED_FORM, n) == fibonacci(LINEAR, n)) {
    fprintf(stderr,
            "closed form is now exact at n=%d; "
            "the documented precision limit changed\n",
            n);
    failures++;
  }

  return failures;
}

int main(void) {
  int failures = 0;
  // begin implementation loop
  for (size_t implementation_index = 0;
       implementation_index <
       sizeof implementations / sizeof implementations[0];
       ++implementation_index) {
    const struct implementation *implementation =
        &implementations[implementation_index];
    // begin test case loop
    for (size_t case_index = 0;
         case_index < sizeof test_cases / sizeof test_cases[0]; ++case_index) {
      const struct test_case *test = &test_cases[case_index];
      int64 got = fibonacci(implementation->mask, test->input);

      if (got != test->expected) {
        fprintf(stderr, "%s: F(%d) returned %lld; expected %lld\n",
                implementation->name, test->input, got, test->expected);
        ++failures;
      }
    }
    // end test case loop
  }
  // end implementation loop

  // begin closed form precision test
  failures += test_closed_form_precision_limit();
  // end closed form precision test

  // report failures (if available)
  if (failures != 0) {
    fprintf(stderr, "%d Fibonacci test(s) failed\n", failures);
    return EXIT_FAILURE;
  }

  puts("All Fibonacci implementations passed 7 test cases.");
  return EXIT_SUCCESS;
}
