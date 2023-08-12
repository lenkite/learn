use std::fs::File;
use std::io::{BufRead, BufReader, Cursor};
use std::io::{self, Read, Write, ErrorKind};

#[test]
#[allow(dead_code)]
#[allow(unused)]
pub fn test_buf_full() {
    let mut buf: [u8; 5] = [0; 5];
    let mut c = Cursor::new(&mut buf[..]);
    // write_cursor(&mut c);
    c.write(&[1; 3][..]).unwrap();
    println!("1. buf: {:?}", c.get_ref());
    let r = c.write(&[2; 5][..]);
    println!("2. buf: {:?}", c.get_ref());
    // for i in 0..10 {
    //     c.write(&[i as u8])?;
    // }
}
#[test]
#[allow(dead_code)]
#[allow(unused)]
pub fn test_read_from_file() -> io::Result<()> {
    let f = File::open("/tmp/foo.txt")?;
    println!("(test_read_from_file) f: {:?}", &f);
    let f = BufReader::new(f);
    // let vec1: Vec<_> = f.lines().collect();
    // println!("vec1: {:?}", vec1);
    let result: Result<Vec<String>, io::Error> = Result::from_iter(f.lines());
    println!("result: {:?}", result);
    Ok(())
}

// fn write_cursor<W: Write>(writer: &mut W)  {
//     writer.write(&[1; 3][..]).unwrap();
//     let r = writer.write(&[2; 5][..]);
//     println!("(write_cursor) r: {:?}", r);
// }

const DEFAULT_BUF_SIZE: usize = 8 * 1024;

pub fn copy<R, W>(reader: &mut R, writer: &mut W) -> io::Result<u64>
    where R: Read + ?Sized, W: Write + ?Sized {
    let mut buf = [0; DEFAULT_BUF_SIZE];
    let mut written = 0;
    loop {
        let len = match reader.read(&mut buf) {
            Ok(0) => return Ok(written),
            Ok(len) => len,
            Err(ref e) if e.kind() == ErrorKind::Interrupted => continue,
            Err(e) => return Err(e),
        };
        writer.write_all(&buf[..len])?;
        written += len as u64;
    }
}


