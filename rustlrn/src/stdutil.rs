#[allow(dead_code)]
pub fn demo_cmp() {
    println!("demo_cmp");
    use std::cmp::{max, min};
    println!("demo_cmp: min: {}", min(3, 4));
    println!("demo_cmp: max: {}", max(3, 4));
}