package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64

	fmt.Println("int size: ", binary.Size(i))
	fmt.Println("int8 size: ", binary.Size(i8))
	fmt.Println("int16 size: ", binary.Size(i16))
	fmt.Println("int32 size: ", binary.Size(i32))
	fmt.Println("int64 size: ", binary.Size(i64))

	var f32 float32
	var f64 float64

	fmt.Println("float32 size: ", binary.Size(f32))
	fmt.Println("float64 size: ", binary.Size(f64))

	var b bool

	fmt.Println("bool size: ", binary.Size(b))

	var s string
	fmt.Println("string size: ", binary.Size(s))

	var sl_int []int64
	sl_int = append(sl_int, 1)
	sl_int = append(sl_int, 2)
	fmt.Println("slice for int64 size: ", binary.Size(sl_int))
}
