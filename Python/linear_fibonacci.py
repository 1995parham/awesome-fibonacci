def fibo(n):
    """Iterate upward, keeping only the last two values.

    Python's integers are arbitrary precision, so this stays exact for any n
    you have the patience to wait for.
    """
    a, b = 0, 1

    for _ in range(n):
        a, b = b, a + b

    return a
