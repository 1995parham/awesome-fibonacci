IDENTITY = ((1, 0), (0, 1))
BASE = ((1, 1), (1, 0))


def _mul(m, o):
    """Multiply two 2x2 matrices given as nested tuples."""
    return (
        (
            m[0][0] * o[0][0] + m[0][1] * o[1][0],
            m[0][0] * o[0][1] + m[0][1] * o[1][1],
        ),
        (
            m[1][0] * o[0][0] + m[1][1] * o[1][0],
            m[1][0] * o[0][1] + m[1][1] * o[1][1],
        ),
    )


def fibo(n):
    """Raise [[1,1],[1,0]] to the n'th power; F(n) lands in the top-right cell.

    Exponentiation by squaring gets there in O(log n) multiplications rather
    than n of them.
    """
    result, base = IDENTITY, BASE

    while n > 0:
        if n & 1:
            result = _mul(result, base)

        base = _mul(base, base)
        n >>= 1

    return result[0][1]
