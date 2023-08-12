
#[allow(dead_code)]
pub fn simple_map_fold() {
    println!("(simple_map_fold)");
    let accum = &[1, 2, 3, 4, 5, 6, 7, 8]
        .iter()
        .map(|x| x + 3)
        .fold(0, |x, y| x + y);
    println!("accum = {}", accum);
}

#[allow(dead_code)]
pub fn demo_people_iter() {
    println!("(demo_people_iter)");
    let pp = People {
        names: vec![
            "bingo".to_string(),
            "tringo".to_string(),
            "hingo".to_string(),
        ],
    };
    for p in &pp {
        println!("(demo_people_iter) p: {}", p);
    }
    for q in &pp {
        println!("(demo_people_iter) q: {}", q);
    }
}

#[allow(dead_code)]
#[derive(Debug)]
pub struct People {
    pub names: Vec<String>,
}

impl<'a> IntoIterator for &'a People {
    type Item = &'a String;
    type IntoIter = PeopleIter<'a>;

    fn into_iter(self) -> Self::IntoIter {
        PeopleIter { names: &self.names }
    }
}

pub struct PeopleIter<'a> {
    names: &'a [String],
}

impl<'a> Iterator for PeopleIter<'a> {
    type Item = &'a String;

    fn next(&mut self) -> Option<Self::Item> {
        match self.names.split_first() {
            Some((next, rest)) => {
                self.names = rest;
                Some(next)
            }
            None => None,
        }
    }
}

#[test]
#[allow(dead_code)]
fn demo_iter() {
    let v = vec![4, 20, 12, 8, 6];
    let mut it1 = v.into_iter();
    assert_eq!(it1.next(), Some(4));
    // assert_eq!(it1.next(), Some(20));
    // assert_eq!(it1.next(), Some(12));
    // assert_eq!(it1.next(), Some(8));
    // assert_eq!(it1.next(), Some(6));
    // assert_eq!(it1.next(), None);

//    let s = "sdfasdadsf";
}

#[test]
#[allow(dead_code)]
fn demo_string_drain() {
    use std::iter::FromIterator;
    let mut outer = "Earth".to_string();
    let drain = outer.drain(..);
    let inner = String::from_iter(drain);
    println!("outer: {}", outer);
    println!("inner: {}", inner);
}

#[test]
#[allow(dead_code)]
fn demo_str_split() {
    let _x1 = &['-', ' ', ':', '@'][..];
    let x2 = &['-', ' ', ':', '@'];
    let v: Vec<&str> = "2020-11-03 23:59".split(x2).collect();
    println!("v: {:?}", v);
    assert_eq!(v, ["2020", "11", "03", "23", "59"]);
}

#[test]
#[allow(dead_code)]
fn demo_inspect() {
    let a = [1, 4, 2, 3];

// this iterator sequence is complex.
    let sum = a.iter()
        .cloned()
        .filter(|x| x % 2 == 0)
        .fold(0, |sum, i| sum + i);

    println!("{sum}");

// let's add some inspect() calls to investigate what's happening
    let sum = a.iter()
        .cloned()
        .inspect(|x| println!("about to filter: {x}"))
        .peekable()
        .filter(|x| x % 2 == 0)
        .inspect(|x| println!("made it through filter: {x}"))
        .fold(0, |sum, i| sum + i);
    println!("{sum}");
}

#[test]
#[allow(dead_code)]
#[allow(unused)]
fn demo_zip() {
    let s1 = &[1, 2, 3];
    let s2 = &['a', 'b', 'c', 'd'];

    let mut iter = s1.iter().zip(s2);
    for e in iter {
        println!("{:?}", e);
    }
}

#[test]
#[allow(dead_code)]
fn demo_by_ref() {
    let message = "To: jimb\r\n\
    From: id\r\n\
\r\n\
Oooooh, donuts!!\r\n";
    let mut lines = message.lines();
    println!("Headers:");
    for header in lines.by_ref().take_while(|l| !l.is_empty()) {
        println!("{}", header);
    }
    println!("\nBody:");
    for body in lines {
        println!("{}", body);
    }
}

#[test]
#[allow(dead_code)]
fn demo_cloned() {
    let a = ['1', '2', '3', '∞'];
    assert_eq!(a.iter().next(), Some(&'1'));
    assert_eq!(a.iter().cloned().next(), Some('1'));
    assert_eq!(a.into_iter().next(), Some('1'));
}

