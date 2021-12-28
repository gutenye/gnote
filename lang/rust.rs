#![allow(unused)]

// line comment
/* block comment */
/// doc line comment, **markdown**
/** doc block comment */

// Error
enum Result<T, E> { Ok(T), Err(E) }
let result = result.unwrap_or_else(|error| {
	if (error.kind() == ErrorKind::NotFound) {
		println!("File not found");
	}
	1
});
let result = match result { Ok(value) => value, Err(error) => panic!() };
let result = result.expect("Error Message");  // panic with message if error
let result = result.unwrap();                 // panic if error
let result = result.unwrap_or(default);       // return default value if error

// Module
use lib1::lib2;
use lib1::{lib2, lib3};
lib2::hello();
mod lib1; // lib1.ts
mod lib1 {
	pub mod lib2 {
		pub fn hello() { }
	}
}

// Print
println!("Hello {} {}", 1, "a");
format!("Hello {} {}", 1, "a");

// Variables
let a: u32 = 1;      // immutable
let mut a = 1;  // mutable
const A = 1;

// Control Flow
if a == 0 { 1 } else { 2 }
let a = if a == 0 { 1 } else { 2 };
match a { 1 => 1, _ => 2 }
for i in 0..10 { i }
while a != 0 { }
loop { break }
let a = loop { break 1 }



