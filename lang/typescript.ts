// line comment
/* block comment */

// Same thing written in multiple ways

// Error
async function propagate_error(): Promise<String> {
	throw new Error("Message1")
	throw new Error("Message2")
	return "ok"
}
try {
	await propagate_error()
} catch (err) {
	console.log(err)
}
propagate_error().catch()

// Module
import lib1, { lib1A, lib1B as lib1C } from './typescript/lib1'
import * as lib1All from './typescript/lib1'
import './typescript/lib1'
export class ExportClass {}
export const exportVar = 1
export function exportFunction() {}
export { lib1A, lib1C as lib1D }
export default 1
export * from './typescript/lib1'
export { default as lib1X, lib1Y, lib1Y as lib1Z } from './typescript/lib1'

// Print & Format
let print1 = "a"
console.log(`Hello ${print1} ${JSON.stringify(print1)} ${JSON.stringify(print1, null, 2)}`) // print string, object, pretty object

// Variables
const var1 = 1; // imutable variable
let var2 = 1;  // mutable variable
const VAR_3 = 1;

// Control Flow
if (var1 === 1) { } else { }
const control1 = var1 === 1 ? 1 : 2
switch (var1) {
	case 1:
		break
	default:
		break
}
for (let i = 0; i < 10; i++) { }
for (const value of [1]) { }
for (const [value, i] of [1].entries()) { }
for (const [key, value] of Object.entries({ a: 1 })) {}
for (const key in { a: 1 }) {}
[1].forEach(value => { })
[1].map(value => value)
while (var1 !== 1) { break; continue; }
while (true) {}

// Function
function fn1(a: string, options: Object = {}, ...rest: any[]): string {
	return "fn1"
}

// Boolean
const bool1: boolean = true

// Integer
const int1: number = 1

// String
const str1: string = "a"
const str2: string = String("a")
const str3: string = "a".toString()

// Array
const arr1: number[] = [1, 2]
const arr2: number[] = new Array()

// Object
const obj1: Object = {
	a: 1,
	get b() { return 1},
	set b(value) {},
	fn() {},
}

// class
class UserBase {}
class User extends UserBase {
	static a = 1
	static staticMethod() {}

	a: number
	b: number = 2

	constructor(a: number) {
		super()
		this.a = a
	}

	method() {
		this.a
	}
}