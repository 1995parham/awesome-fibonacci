use fibonacci::{
    CLOSED_FORM_EXACT_UP_TO, closed_form, fast_doubling, linear, matrix, memoized, recursive,
};

/// (n, F(n)) with the standard indexing: F(0) = 0, F(1) = 1.
const CASES: [(u64, u64); 13] = [
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
];

type Impl = (&'static str, fn(u64) -> u64);

const ALL: [Impl; 6] = [
    ("recursive", recursive),
    ("memoized", memoized),
    ("linear", linear),
    ("matrix", matrix),
    ("fast_doubling", fast_doubling),
    ("closed_form", closed_form),
];

#[test]
fn known_values() {
    for (name, fibo) in ALL {
        for (n, want) in CASES {
            assert_eq!(fibo(n), want, "{name}({n})");
        }
    }
}

/// The implementations against each other over a contiguous range the table
/// only samples. Bounded by recursive's exponential cost.
#[test]
fn agreement() {
    for n in 0..30 {
        let want = linear(n);
        for (name, fibo) in ALL {
            assert_eq!(fibo(n), want, "{name}({n})");
        }
    }
}

/// F(90) is close to the u64 ceiling (F(93) is the last value that fits) and
/// far past anything the closed form can reach. The exact implementations must
/// still agree there. Recursive is excluded as too slow and closed_form as
/// no longer exact.
#[test]
fn large_n() {
    let want = linear(90);
    assert_eq!(want, 2880067194370816120);

    for (name, fibo) in ALL {
        if matches!(name, "recursive" | "closed_form") {
            continue;
        }
        assert_eq!(fibo(90), want, "{name}(90)");
    }
}

/// Where Binet's formula stops being exact -- a documented property, but a
/// silent change to it should still fail the build.
#[test]
fn closed_form_precision_limit() {
    for n in 0..=CLOSED_FORM_EXACT_UP_TO {
        assert_eq!(
            closed_form(n),
            linear(n),
            "closed form should be exact at {n}"
        );
    }

    let n = CLOSED_FORM_EXACT_UP_TO + 1;
    assert_ne!(
        closed_form(n),
        linear(n),
        "closed form is now exact at {n}; the documented limit moved"
    );
}
