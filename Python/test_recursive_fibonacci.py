import unittest

from recursive_fibonacci import fibo

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


class TestRecursive(unittest.TestCase):
    def test_known_values(self):
        for n, want in CASES:
            with self.subTest(n=n):
                self.assertEqual(fibo(n), want)


if __name__ == "__main__":
    unittest.main()
