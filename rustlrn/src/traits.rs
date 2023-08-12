use crate::structs::*;
use std::borrow::Cow;
use std::collections::HashSet;
use std::ops::Add;

impl std::ops::Neg for Number {
    type Output = Number;

    fn neg(self) -> Self {
        Self {
            value: -self.value,
            odd: self.odd,
        }
    }
}

impl std::ops::Add for Number {
    type Output = Number;
    fn add(self, rhs: Self) -> Self::Output {
        Self {
            value: self.value + rhs.value,
            odd: self.odd,
        }
    }
}

#[allow(dead_code)]
pub fn demo_neg_trait() {
    let m = Number {
        odd: true,
        value: 987,
    };
    println!("(demo_neg_trait) m={:?}", m); // prints "-987"
    let n = -m; // this is only possible because we implemented `Neg`
    println!("(demo_neg_trait) n={:?}", n); // prints "-987"
}

#[allow(dead_code)]
pub fn sum<A: Add>(a: A, b: A) -> impl Add
where
    <A as Add>::Output: Add,
{
    a + b
}

pub trait Float {
    const ZERO: Self;
    const ONE: Self;
}

impl Float for f32 {
    const ZERO: f32 = 0.0;
    const ONE: f32 = 1.0;
}

impl Float for f64 {
    const ZERO: f64 = 0.0;
    const ONE: f64 = 1.0;
}

#[allow(dead_code)]
pub fn add_one<T>(value: T) -> T
where
    T: Add<Output = T> + Float,
{
    value + T::ONE
}

#[allow(dead_code)]
pub fn add1<T>(v1: T, v2: T) -> T
where
    T: Add<Output = T>,
{
    v1 + v2
}

#[derive(Debug, Copy, Clone, PartialEq)]
pub struct Complex<T> {
    re: T,
    im: T,
}

// impl<T> Add for Complex<T> where
//     T: Add<Output=T>, {
//     type Output = Self;
//     fn add(self, rhs: Self) -> Self {
//         Complex {
//             re: self.re + rhs.re,
//             im: self.im + rhs.im,
//         }
//     }
// }

impl<L, R> Add<Complex<R>> for Complex<L>
where
    L: Add<R>,
{
    type Output = Complex<L::Output>;
    fn add(self, rhs: Complex<R>) -> Self::Output {
        Complex {
            re: self.re + rhs.re,
            im: self.im + rhs.im,
        }
    }
}

use std::ops::Neg;

impl<T> Neg for Complex<T>
where
    T: Neg<Output = T>,
{
    type Output = Complex<T>;
    fn neg(self) -> Complex<T> {
        Complex {
            re: -self.re,
            im: -self.im,
        }
    }
}

use std::ops::BitAnd;

impl<T> BitAnd for Complex<T>
where
    T: BitAnd<Output = T>,
{
    type Output = Self;
    fn bitand(self, rhs: Self) -> Self::Output {
        Complex {
            re: self.re & rhs.re,
            im: self.im & rhs.im,
        }
    }
}

use std::ops::AddAssign;

impl<T> AddAssign for Complex<T>
where
    T: AddAssign<T>,
{
    fn add_assign(&mut self, rhs: Self) {
        self.re += rhs.re;
        self.im += rhs.im;
    }
}

#[test]
fn demo_complex_bitand() {
    let c1 = Complex { re: 1, im: 3 };
    let c2 = Complex { re: 0, im: 1 };
    println!("c1 = {:?}, c2={:?}", c1, c2);
    println!("c1 & c2 = {:?}", c1 & c2);
}

/// .
#[test]
fn demo_complex_mult() {
    let mut x = Complex { re: 5, im: 2 };
    let y = Complex { re: 2, im: 5 };
    x += y;
    println!("(demo_complex_mult) x = {:?}", x);
    assert_eq!(x, Complex { re: 7, im: 7 });
}

#[test]
fn sort_demo() {
    use std::cmp::Reverse;
    let mut v = vec![1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 3, 3, 2, 2, 1, 0];
    v.sort_by_key(|&num| {
        let x = (num > 3, Reverse(num));
        println!("num: {}, x: {:?}", num, x);
        x
    });
    // v.sort_by_key(|&num| Reverse(num));
    println!("(sort_by_key) v = {:?}", v);
    // assert_eq!(v, vec![3, 2, 1, 6, 5, 4]);
    // num>2 (sort_by_key) v = [2, 1, 6, 5, 4, 3]
    println!("{}", Reverse(1) < Reverse(2));
    println!("{}", false < true);
}

#[test]
fn demo_complex_neg() {
    let c = Complex { re: 1.0, im: 1.0 };
    println!("c = {:?}", c);
    println!("-c = {:?}", -c);
}

#[allow(dead_code)]
struct Data<'a> {
    header: [u8; 4],
    body: &'a [u8],
}

#[allow(dead_code)]
fn destructure(input: &[u8]) -> Data {
    match input {
        [0x04, 0x07, a, b, c, d, body @ ..] => Data {
            header: [*a, *b, *c, *d],
            body,
        },
        [_a, _b, _, _, _, _, ..] => panic!("invalid magic number"),
        [0x04, 0x07, ..] => panic!("insufficient packet length"),
        [_a, _b, ..] => panic!("invalid magic number, and insufficient packet length"),
        [..] => panic!("insufficient packet length"),
    }
}

#[allow(dead_code)]
#[test]
fn test_to_owned() {
    let a1: &[&str] = &["a", "b", "c"];
    let b1 = a1.to_owned();
    let c1 = Cow::from(a1);
    let mut d1 = c1.into_owned();
    d1.push("bingo");

    println!("a1: {:?}, b1: {:?}, d1: {:?}", a1, b1, d1);
    //    println!("a1: {:?}, b1: {:?}, c1: {:?}, d1: {:?}", a1, b1, c1, d1);

    let a2: &[i32] = &[1, 2, 3];
    let b2 = a2.to_owned();
    println!("a2: {:?}, b2: {:?}", a2, b2);
}
//use std::collections::HashSet;

#[derive(Clone, Debug)]
#[allow(dead_code)]
struct Element {
    id: usize,
}

#[allow(dead_code)]
fn get_unique_cow(input: &[Element]) -> Cow<[Element]> {
    let mut set = HashSet::new();
    let mut contains_duplicate = false;
    for element in input {
        if set.contains(&element.id) {
            contains_duplicate = true;
        }
        set.insert(element.id);
    }
    if !contains_duplicate {
        return Cow::Borrowed(input);
    }
    let mut ret = Vec::new();
    for element in input {
        if set.contains(&element.id) {
            ret.push(element.to_owned());
            set.remove(&element.id);
        }
        // duplicate
    }
    return Cow::Owned(ret);
}

#[test]
#[allow(dead_code)]
fn demo_cow() {
    let e1 = vec![
        Element { id: 1 },
        Element { id: 1 },
        Element { id: 2 },
        Element { id: 2 },
        Element { id: 3 },
    ];
    let s = "bingo";
    let _t1 = s.to_owned();
    let _t2 = s.to_string();
    let e2 = get_unique_cow(&e1);
    println!("e2: {:?}", e2);
}
#[test]
#[allow(dead_code)]
#[allow(unused)]
fn demo_iter_any() {
    let vec1 = vec![1, 2, 3];
    let vec2 = vec![4, 5, 6];

    // `iter()` for vecs yields `&i32`. Destructure to `i32`.
    println!("A> 2 in vec1: {}", vec1.iter().any(|&x| x == 2));

    println!("B> in vec1: {}", vec1.iter().any(|x| *x == 2));
    println!("C> vec1: {:?}", vec1);
}
