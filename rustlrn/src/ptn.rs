#[allow(dead_code)]
use crate::structs::*;
#[allow(dead_code)]
pub fn demo_let_ptn1() {
    println!("demo_let_ptn1");
    let one = Number {
        odd: true,
        value: 1,
    };
    let two = Number {
        odd: false,
        value: 2,
    };
    print_number(one);
    print_number(two);
}

#[allow(dead_code)]
fn print_number(n: Number) {
    if let Number { odd: true, value } = n {
        println!("Odd number: {}", value);
    } else if let Number {
        odd: false,
        value: val,
    } = n
    {
        println!("Even number: {}", val);
    }
}

#[allow(dead_code)]
pub fn demo_destructure_update() {
    println!("demo_destructure_update");
    let p = Point { x: 1.0, y: 2.0 };
    let Point { x, y } = p;
    println!("demo_destructure_update:  (x, y): ({}, {})", x, y);
    let p1 = Point { x: 3.0, ..p };
    let p2 = Point { ..p };
    println!("demo_destructure_update:  p1: {:?}, p2: {:?}", p1, p2);
    let Point { x: x1, .. } = p1;
    println!(
        "demo_destructure_update:  x1: {} after destructure from p1: {:?}",
        x1, p1
    );
    let Point { y, .. } = p2;
    println!(
        "demo_destructure_update:   y: {} after destructure from p2: {:?}",
        y, p2
    );
}
