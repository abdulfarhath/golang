package main

import (
	"fmt"
)

// 1. Built-in Data Types

// Boolean
var b bool = true

// Numeric Types
var i int = 42
var i8 int8 = -127
var i16 int16 = 32767
var i32 int32 = -2147483648
var i64 int64 = 9223372036854775807

var u uint = 42
var u8 uint8 = 255
var u16 uint16 = 65535
var u32 uint32 = 4294967295
var u64 uint64 = 18446744073709551615

var f32 float32 = 3.14
var f64 float64 = 2.7182818284

var c64 complex64 = 1 + 2i
var c128 complex128 = 2 + 3i

// String
var s string = "hello"

// Byte and Rune
var by byte = 'a'      // alias for uint8
var ru rune = '世'     // alias for int32

// 2. Aggregate Types

// Array
var arr [3]int = [3]int{1, 2, 3}

// Slice
var slc []string = []string{"a", "b", "c"}

// Struct
type Person struct {
	Name string
	Age  int
}
var p Person = Person{"Alice", 30}

// Map
var m map[string]int = map[string]int{"a": 1, "b": 2}

// 3. Reference Types

// Pointer
var ptr *int = &i

// Function
func add(a int, b int) int {
	return a + b
}
var fn func(int, int) int = add

// Interface
type Describer interface {
	Describe() string
}

// Channel
var ch chan int = make(chan int)

// 4. User-defined Types

// Type Alias
type MyInt = int

// New Type
type Score int

// Enum-like using iota
type Day int
const (
	Sunday Day = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// Ways to define a datatype in Go:

// 1. Using built-in types directly (e.g., int, string)
// 2. Using type keyword to define a new type
type MyString string

// 3. Using type keyword for struct, interface, function, etc.
type Rectangle struct {
	Width, Height float64
}
type Shape interface {
	Area() float64
}
type Operation func(int, int) int

// 4. Type aliasing
type Age = int

func main() {
	fmt.Println("bool:", b)
	fmt.Println("int:", i, "int8:", i8, "int16:", i16, "int32:", i32, "int64:", i64)
	fmt.Println("uint:", u, "uint8:", u8, "uint16:", u16, "uint32:", u32, "uint64:", u64)
	fmt.Println("float32:", f32, "float64:", f64)
	fmt.Println("complex64:", c64, "complex128:", c128)
	fmt.Println("string:", s)
	fmt.Println("byte:", by, "rune:", ru)
	fmt.Println("array:", arr)
	fmt.Println("slice:", slc)
	fmt.Println("struct:", p)
	fmt.Println("map:", m)
	fmt.Println("pointer:", ptr)
	fmt.Println("function:", fn(2, 3))
	fmt.Println("channel:", ch)
	fmt.Println("type alias MyInt:", MyInt(10))
	fmt.Println("new type Score:", Score(99))
	fmt.Println("enum-like Day:", Monday)
}