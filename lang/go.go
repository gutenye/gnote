
package main

/* Dislike
- no allow unused flag
- no super map func: b := []string{}; for _, v := range a { b = append(b, v) }
- no super sort, join func: sort.{Strings,Ints,..}, strings.Join
- filepath.Glob not support `**`
*/

// Easy to Read
v1 := "a"              // var v = "a"
_, err := f()          // blank identifier
if v2, ok := map["a"]; ok {}     // comma ok idiom, hasKey?
if v := f(); v == 0 {} // v := f(); if v == 0 {}
switch { case a == 0:   }
func f(a, b int) {}    // func f(a int, b int)
const ( _ = iota; KB = 1 << (10 * iota); MB ) // enum

// Allow unused
import _ "fmt"; _ = a   