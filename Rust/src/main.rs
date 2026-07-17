use std::io;

fn main() {
    let mut n = String::new();

    println!("Enter length of fibonacci sequence:");

    io::stdin().read_line(&mut n).expect("Failed to read line");

    let n: u32 = n.trim().parse().expect("Not a number");

    for i in 0..n {
        let x = fibonacci(i);
        println!("{}", x);
    }
}

// returns n'th element of fibonacci sequence in a recursive fashion
fn fibonacci(n: u32) -> u32 {
    if n == 0 {
        0
    } else if n == 1 {
        1
    } else {
        fibonacci(n - 1) + fibonacci(n - 2)
    }
}

#[cfg(test)]
mod tests {
    use super::fibonacci;

    // (n, F(n)) with the standard indexing: F(0) = 0, F(1) = 1.
    // Kept at or below F(47), the last value that fits in a u32.
    const CASES: [(u32, u32); 13] = [
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

    #[test]
    fn known_values() {
        for (n, want) in CASES {
            assert_eq!(fibonacci(n), want, "fibonacci({n})");
        }
    }
}
