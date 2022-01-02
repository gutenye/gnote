package main
import _ "fmt"; _ = a    // allow unused

// Easy to Read
v1 := "a"              // var v = "a"
_, err := f()          // blank identifier
if v2, ok := map["a"]; ok {}     // comma ok idiom, hasKey?
if v := f(); v == 0 {} // v := f(); if v == 0 {}
switch { case a == 0:   }
func f(a, b int) {}    // func f(a int, b int)
const ( _ = iota; KB = 1 << (10 * iota); MB ) // enum
