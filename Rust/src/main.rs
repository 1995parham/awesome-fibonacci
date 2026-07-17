use std::io;

use fibonacci::linear;

fn main() {
    let mut n = String::new();

    println!("Enter length of fibonacci sequence:");

    io::stdin().read_line(&mut n).expect("Failed to read line");

    let n: u64 = n.trim().parse().expect("Not a number");

    for i in 0..n {
        println!("{}", linear(i));
    }
}
