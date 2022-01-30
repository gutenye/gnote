package main

// no allow unused var
// no super map
a := []int{10, 11}; b := []int{}; for _, v := range a { b = append(b, v + 1) } // no super map
// no super sort, join func
sort.{Strings,Ints,..}, strings.Join
// filepath.Glob not support `**`





// Easy to Read
v1 := "a"              // var v = "a"
_, err := f()          // blank identifier
if v2, ok := map["a"]; ok {}     // comma ok idiom, hasKey?
if v := f(); v == 0 {} // v := f(); if v == 0 {}
switch { case a == 0:   }
func f(a, b int) {}    // func f(a int, b int)
const ( _ = iota; KB = 1 << (10 * iota); MB ) // enum
import ( "fmt"; "os" )   // import "fmt"; import "os";

// line comment
/* block comment */

// Allow unused
import _ "fmt"; _ = a   

// Error
func hello() (string, error) {
	return "", errors.New("Error")
	return "Hello", nil

value, err := hello()
value2, err := hello() // special, can assign same err variale multiple times
if err != nil {}

// Module
// root/*.go
package root
// root/dir/*.go
package dir
// a.go
import "example.com/root"
import . "example.com/root"
import root2 "example.com/root"
import "example.com/root/dir"

// Print & Format
println("hello")
fmt.Printf("%v\n", [1])
fmt.Sprintf("%s", "a")

// Variables
var a int
var a int = 1
var a = 1
a := 1

// Control Flow
if a == 0 {} else if {} else {}
if err := f(); err != nil {}
switch a { case 1, 2: fallthrough; default: }
switch { case a == 0 }
switch t := interface.(type) { case int } // special
switch a := f(); a {}
for k, v := range x {}
for k := range x {}
for _, v := range x  {}
for i := 0; i < 10; i++ {}
for i := 0; i < 10 {}
for i < 10 {}
for {}
for i, j := 1, 2; i< 10; i, j = i+1, j+1 { .. }
break [label]; continue [label]; goto [label]
Label: for { break Label }

// Function
func fn1(a, b int, cb func(int) int), args ...interface{}) (int, error) {
	file, err := os.Open("file")
	defer file.Close()  // run at the end of the function
	return x, y
}
func fn1() (x, y int) { x = 1; y = 2; return }
var fn1 func(int) int = func(a int) int { return a }
a, b := fn1(args...)

// Boolean
var a bool = true

// Integer
var a int = 1

// String
var a string = "a"

// Array
var a []int = []int{1,2}
a := []int{10, 11}
// Iterator
for _, v := range a {} 
for i, v := range a {} 
b := []int{}; for _, v := range a { b = append(b, v + 1) } // no super map

// Struct
type Point struct {
type Point struct {
	x int `tag:"json" b:2`
	x, y int
  User
	pkg.User
}
func New(x, y int) *Point {
	return &Point{x, y, User{}}
}
func (p *Point) Method() {}
var p struct {x int}

Point{1,2} //-> Point 
Point{x:1} 
Point{}
New(1,2)
new(Point) //-> *Point
&Point{} //-> *Point
make(T, ...args)


