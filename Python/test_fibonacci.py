import unittest

import closed_form_fibonacci
import fast_doubling_fibonacci
import linear_fibonacci
import matrix_fibonacci
import memoized_fibonacci
import recursive_fibonacci

# (n, F(n)) with the standard indexing: F(0) = 0, F(1) = 1.
CASES = [
    (0, 0),
    (1, 1),
    (2, 1),
    (3, 2),
    (4, 3),
    (5, 5),
    (6, 8),
    (7, 13),
    (8, 21),
    (9, 34),
    (10, 55),
    (20, 6765),
    (30, 832040),
]

IMPLEMENTATIONS = [
    ("recursive", recursive_fibonacci.fibo),
    ("memoized", memoized_fibonacci.fibo),
    ("linear", linear_fibonacci.fibo),
    ("matrix", matrix_fibonacci.fibo),
    ("fast_doubling", fast_doubling_fibonacci.fibo),
    ("closed_form", closed_form_fibonacci.fibo),
]


class TestImplementations(unittest.TestCase):
    def test_known_values(self):
        """Every implementation against the same table."""
        for name, fibo in IMPLEMENTATIONS:
            for n, want in CASES:
                with self.subTest(implementation=name, n=n):
                    self.assertEqual(fibo(n), want)

    def test_agreement(self):
        """The implementations against each other, over a contiguous range.

        Bounded by recursive, which is exponential and gets slow quickly.
        """
        for n in range(30):
            want = linear_fibonacci.fibo(n)
            for name, fibo in IMPLEMENTATIONS:
                with self.subTest(implementation=name, n=n):
                    self.assertEqual(fibo(n), want)


class TestExactImplementations(unittest.TestCase):
    """Python's integers are unbounded, so every implementation except the
    closed form should stay exact arbitrarily far out."""

    EXACT = [impl for impl in IMPLEMENTATIONS if impl[0] not in ("recursive", "closed_form")]

    def test_large_n(self):
        # F(500), which is 105 digits long and would overflow any fixed-width
        # integer type by a wide margin.
        want = linear_fibonacci.fibo(500)
        self.assertEqual(len(str(want)), 105)

        for name, fibo in self.EXACT:
            with self.subTest(implementation=name):
                self.assertEqual(fibo(500), want)


class TestClosedFormPrecisionLimit(unittest.TestCase):
    """Where Binet's formula stops being exact. A documented property rather
    than a bug, but a silent change to it should still fail the build."""

    def test_exact_up_to_limit(self):
        for n in range(closed_form_fibonacci.EXACT_UP_TO + 1):
            with self.subTest(n=n):
                self.assertEqual(closed_form_fibonacci.fibo(n), linear_fibonacci.fibo(n))

    def test_diverges_past_limit(self):
        n = closed_form_fibonacci.EXACT_UP_TO + 1
        self.assertNotEqual(
            closed_form_fibonacci.fibo(n),
            linear_fibonacci.fibo(n),
            f"closed form is now exact at n={n}; the documented limit moved",
        )


if __name__ == "__main__":
    unittest.main()
