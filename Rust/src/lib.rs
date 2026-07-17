//! The Fibonacci sequence, by every algorithm.
//!
//! Every function here returns the same thing for the same input, using the
//! standard indexing: `F(0) = 0`, `F(1) = 1`. They differ only in how they get
//! there, which is the entire point of having six of them.
//!
//! All of them work in `u64`, so `F(93)` is the last value that fits. Several
//! overflow earlier than that internally, and in a debug build (which is what
//! `cargo test` uses) an overflow panics rather than wrapping.

use std::collections::HashMap;

/// Transcribe the definition directly. `O(phi^n)`.
///
/// Every call re-derives both of its subtrees from scratch, so the same `F(k)`
/// is computed over and over. It is here as the baseline the others improve on.
pub fn recursive(n: u64) -> u64 {
    if n == 0 {
        0
    } else if n == 1 {
        1
    } else {
        recursive(n - 1) + recursive(n - 2)
    }
}

/// The recursive version with each `F(n)` remembered the first time it is
/// computed. `O(n)`.
///
/// The cache is what collapses the exponential call tree into n additions.
pub fn memoized(n: u64) -> u64 {
    fn go(n: u64, cache: &mut HashMap<u64, u64>) -> u64 {
        if let Some(&f) = cache.get(&n) {
            return f;
        }

        let f = go(n - 1, cache) + go(n - 2, cache);
        cache.insert(n, f);

        f
    }

    go(n, &mut HashMap::from([(0, 0), (1, 1)]))
}

/// Iterate upward keeping only the last two values. `O(n)` time, `O(1)` space.
pub fn linear(n: u64) -> u64 {
    let (mut a, mut b) = (0u64, 1u64);

    for _ in 0..n {
        (a, b) = (b, a + b);
    }

    a
}

/// Multiply two 2x2 matrices laid out in row-major order.
fn mul(m: [[u64; 2]; 2], o: [[u64; 2]; 2]) -> [[u64; 2]; 2] {
    [
        [
            m[0][0] * o[0][0] + m[0][1] * o[1][0],
            m[0][0] * o[0][1] + m[0][1] * o[1][1],
        ],
        [
            m[1][0] * o[0][0] + m[1][1] * o[1][0],
            m[1][0] * o[0][1] + m[1][1] * o[1][1],
        ],
    ]
}

/// Raise `[[1,1],[1,0]]` to the n'th power; `F(n)` lands in the top-right cell.
/// `O(log n)` via exponentiation by squaring.
pub fn matrix(n: u64) -> u64 {
    let mut result = [[1, 0], [0, 1]];
    let mut base = [[1, 1], [1, 0]];
    let mut e = n;

    while e > 0 {
        if e & 1 == 1 {
            result = mul(result, base);
        }

        e >>= 1;

        // Squaring after the last set bit is wasted work, and worse than
        // wasted here: that final square reaches F(128), which overflows u64
        // even when the result F(n) itself fits.
        if e > 0 {
            base = mul(base, base);
        }
    }

    result[0][1]
}

/// Return `(F(n), F(n+1))`, halving n on each step.
///
/// The two identities that make this work:
///
/// ```text
/// F(2k)   = F(k) * (2*F(k+1) - F(k))
/// F(2k+1) = F(k)^2 + F(k+1)^2
/// ```
fn pair(n: u64) -> (u64, u64) {
    if n == 0 {
        return (0, 1);
    }

    let (a, b) = pair(n >> 1);
    let c = a * (2 * b - a);
    let d = a * a + b * b;

    if n & 1 == 1 { (d, c + d) } else { (c, d) }
}

/// The matrix method with the matrix algebra folded out by hand, so it is also
/// `O(log n)` but does fewer multiplications.
pub fn fast_doubling(n: u64) -> u64 {
    pair(n).0
}

/// Largest n for which [`closed_form`] still rounds to the right integer.
///
/// Measured, not derived. Go's `f64` closed form diverges at the same n=76,
/// while Python's `**` gives out earlier at n=71: the limit is a property of
/// how each language rounds `pow`, not of the mathematics.
pub const CLOSED_FORM_EXACT_UP_TO: u64 = 75;

/// Binet's formula: `F(n) = (phi^n - psi^n) / sqrt(5)`.
///
/// The only implementation here that does not compute `F(n)` exactly. It rounds
/// an `f64`, and the rounding error compounds with n until it exceeds 1/2 and
/// lands on the wrong integer -- see [`CLOSED_FORM_EXACT_UP_TO`].
pub fn closed_form(n: u64) -> u64 {
    let sqrt5 = 5f64.sqrt();
    let phi = (1.0 + sqrt5) / 2.0;
    let psi = (1.0 - sqrt5) / 2.0;

    ((phi.powi(n as i32) - psi.powi(n as i32)) / sqrt5).round() as u64
}
