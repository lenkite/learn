
#[test]
#[allow(dead_code)]
#[allow(unused)]
fn slice_split() {
    let v = vec![11, 22, 33, 0, 44, 55, 0, 0, 66];
    let splits = v.split(|&x| x == 0);
    for s in splits {
        println!("s = {:?}", s);
    }
}

