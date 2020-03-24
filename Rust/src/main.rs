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

