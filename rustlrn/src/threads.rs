use std::thread;
use std::sync::mpsc::channel;
use byteorder::LittleEndian;

#[test]
#[allow(dead_code)]
#[allow(unused)]
fn test_channel_mpsc() {
    let (tx, rx) = channel();
    for i in 0..10 {
        let tx = tx.clone();
        thread::spawn(move || {
            tx.send(i).unwrap();
            println!("(test_channel_mpsc) Sent i: {}", i);
            ;
        });
    }

    for _ in 0..10 {
        let j = rx.recv().unwrap();
        println!("(test_channel_mpsc) received j: {}", j);
        ;
        assert!(0 <= j && j < 10);
    }
}

#[test]
#[allow(dead_code)]
#[allow(unused)]
fn test_endian() {
    use std::io::Cursor;
    use byteorder::{BigEndian, ReadBytesExt};

    let v = vec![2u8, 5u8, 3u8, 0u8];
    let mut rdr = Cursor::new(v);

    //    vector elements:        0x02, 0x05, 0x03, 0x00


    // 0x1234 5678  in little endian is:
    // 0x7856 3412


    // let a = rdr.read_u16::<BigEndian>().unwrap();
    // println!("a = {} or {:#02x}", a, a);
    // let b = rdr.read_u16::<BigEndian>().unwrap();
    // println!("b = {} or {:#02x}", b, b);

    let a = rdr.read_u16::<LittleEndian>().unwrap();
    println!("a = {} or {:#02x}", a, a);
    let b = rdr.read_u16::<LittleEndian>().unwrap();
    println!("b = {} or {:#02x}", b, b);
    /*
    a = 517 or 0x205
    b = 768 or 0x300
     */
    // let c = rdr.read_u16::<BigEndian>();
    // println!("c = {}", c.unwrap());
}
