def _pair(n):
    """Return (F(n), F(n+1)), halving n on each step.

    The two identities that make this work:

        F(2k)   = F(k) * (2*F(k+1) - F(k))
        F(2k+1) = F(k)^2 + F(k+1)^2

    They are the matrix method with the matrix algebra worked out by hand,
    which is why this is also O(log n) but does less multiplying.
    """
    if n == 0:
        return 0, 1

    a, b = _pair(n >> 1)
    c = a * (2 * b - a)
    d = a * a + b * b

    if n & 1:
        return d, c + d

    return c, d


def fibo(n):
    return _pair(n)[0]
