#[allow(dead_code)]
#[derive(Debug, Copy, Clone)]
pub struct Point {
    pub x: f64,
    pub y: f64,
}

#[allow(dead_code)]
#[derive(Debug)]
pub struct Number {
    pub odd: bool,
    pub value: i32,
}



impl Number {
    fn is_strictly_positive(self) -> bool {
        self.value > 0
    }
}

#[allow(dead_code)]
pub fn demo_method() {
    let minus_two = Number {
        odd: false,
        value: -2,
    };
    let one = Number {
        odd: true,
        value: 1,
    };
    println!("(demo_method) positive? {}", minus_two.is_strictly_positive());
    println!("(demo_method) positive? {}", one.is_strictly_positive());
    // this prints "positive? false"
}
