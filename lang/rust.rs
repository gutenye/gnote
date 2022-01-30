// allow unused var
#![allow(unused)]

// map
let a = [10, 11].iter(); let b: Vec<_> = a.map(|v| v + 1).collect();

// Same thing written in multiple ways
match result { Ok(v) => v, Err(err) => panic!() }
result.expect("Error");
result.unwrap();
result.unwrap_or(default)
result.unwrap_or_else(|error| { default })
//
let a: &str = "a";
let a: String = String::from("a");
let a: String = "a".to_string();
let a: String = "a".into();

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
resut.is_err();
fn propagate_error() -> Result<String, io::Error> {
	let result = File::open()?;
	let result = File::open()?.read_to_string()?;
	let result = match (File::open()) {
		Ok(value) => value,
		Err(error) => return Err(error)
	}
	Ok(result)
}

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

// Print & Format
println!("Hello {} {:?} {:#?}", 1, "a"); // print string, object, pretty object
format!("Hello {} {}", 1, "a");

// Variables
let a: u32 = 1;  // immutable variable
let mut a = 1;   // mutable variable
const A = 1;

// Control Flow
if a == 0 { 1 } else { 2 }
let a = if a == 0 { 1 } else { 2 };
match a { 1 => 1, _ => 2 }
for i in 0..10 { i }
while a != 0 { }
loop { break; continue; }
let a = loop { break 1 }

// Function
fn fn1(a: u32) -> u32 {
	return 1;
	1
}

// Macro


// Boolean
let a: bool = true;

// Integer
let a: u32 = 1;

// String
let a: &str = "a";
let a: &str = &String::from("a");
let a: String = "a".to_string();
let a: String = "a".into();

// Array
let a: [u32; 2] = [1, 2];
let a: [u32; 2] = [1; 2];
let a: Vec<u32> = vec![1, 2];
let a: Vec<u32> = Vec::new();
// Iterator
let a = [10, 11].iter();
for v in a {}
for (i, v) in a.enumerate() {}
let b: Vec<_> = a.map(|v| v + 1).collect();

// Struct
struct Black;         // unit struct
struct Point(x, y);   // tuple struct
struct User {         // object struct
	name: String,
	active: bool,
}
impl User {
	fn new(name: String, active: bool) -> User {
		User { name, active }
	}
	fn method(&self) {
		println!("Hello {}", self.name);
	}
}
let point: Point = Point(1, 2);
let user: User = User {
	name: "Name",
	active: true,
	..user1,
}
let user: User = User::new("Name".to_string(), false);
user.name
user.method()

// Enum
enum Enum {
	A,
	A(User, u32),
	A { x: u32 },
}
impl Enum {
	fn method(&self) { }
}
let a: Enum = Enum::A;
let a: Enum = Enum::A(user, 1);
let a: Enum = Enum::A { x: 1 };
match a {
	Enum::A => 1,
	Enum::A(user, 1) => user.name,
	Enum::A { x } => x,
}
