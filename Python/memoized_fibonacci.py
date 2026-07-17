def fibo(n, cache=None):
    """Recursive fibonacci, but each F(n) is computed only once.

    The cache is what separates this from recursive_fibonacci: it turns an
    exponential call tree into n additions. functools.cache as a decorator
    would be the idiomatic Python way to say this, but spelling the cache out
    keeps the algorithm visible.
    """
    if cache is None:
        cache = {0: 0, 1: 1}

    if n in cache:
        return cache[n]

    cache[n] = fibo(n - 2, cache) + fibo(n - 1, cache)

    return cache[n]
