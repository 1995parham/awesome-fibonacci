fn main() {
    const N: u32 = 10;
    for n in 0..N {
        let x = fibonacci(n);
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
        fibonacci(n-1) + fibonacci(n-2)
    }
}
