extern crate core;

mod iter;
mod ptn;
mod stdutil;
mod structs;
mod traits;
mod coll;
mod format;
mod fnpointers;
mod iolrn;
mod threads;

fn main() {
    // iter::simple_map_fold();
    //stdutil::demo_cmp();
    //structs::demo_destructure_update();
    // ptn::demo_let_ptn1();
    // structs::demo_method();
    // traits::demo_neg_trait();
    // iter::demo_people_iter();
    // let v = vec![0, 1, 2];
    // let iter1 = v.iter();
    // let s = traits::sum(1, 2);
    let s = traits::add_one(3.0);
    println!("s:{}", s);
}
