#[test]
#[allow(dead_code)]
#[allow(unused)]
fn demo_fnpoint1() {
    fn sum(x: i32, y: i32) -> i32 {
        x + y
    }
    // Explicit coercion to `fn` type is required...
    let op_a: fn(i32, i32) -> i32 = sum;
    let op_b = sum;
    let op_c = sum;

    assert!(op_a == op_b);
    // assert!(op_b  == op_c); //doesn't work.
    // assert_eq!(opA, opB);
}
