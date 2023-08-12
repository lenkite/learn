#[test]
#[allow(dead_code)]
#[allow(unused)]
fn test_format1() {
    println!("1.>>{:12}<<", "bookends");
    // println!("2.>>{:20} {:04x} {:10x}<<", "adc #42", 105, 42);
    // println!("3.>>{:20} {:x} {:02x}<<", "adc #42", 105, 42);
    println!("2.>>{:9x}|{:3x}|{:05x}<<", 105, 105, 105);
    // println!("3>>{1:02x} {2:02x} {0}", "");
    println!("3.>>{lsb:02x} {msb:02x} {insn}<<", insn="adc #42", lsb=105, msb=42);
    println!("4.>>{:=^12}<<", "bookends");
    println!("5.>>{:#^12}<<", "bookends");
    println!("6.>>{:1^12}<<", "bookends");
    println!("7.>>{:12}<<", 999);
}
