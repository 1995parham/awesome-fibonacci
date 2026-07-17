import math

SQRT5 = math.sqrt(5)
PHI = (1 + SQRT5) / 2
PSI = (1 - SQRT5) / 2

#: Largest n for which this still rounds to the right integer. Beyond it the
#: rounding error in ** has compounded past 1/2 and lands on a neighbour.
#: Measured, not derived. Go's closed form survives to 75 doing the same
#: arithmetic, because math.Pow and ** round differently.
EXACT_UP_TO = 70


def fibo(n):
    """Binet's formula: F(n) = (phi**n - psi**n) / sqrt(5).

    The only implementation here that does not compute F(n) exactly. Python's
    integers are unbounded but its floats are ordinary doubles, so this one
    inherits the double's precision rather than the integer's range, and it is
    only correct up to EXACT_UP_TO.
    """
    return round((PHI**n - PSI**n) / SQRT5)
