package main

import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

/*
使用binary.Size, 必须是固定长度的数据类型，
比如int8, int16, int32, int64, float32, float64, bool类型。
对于依赖操作系统而变长的数据类型，返回-1，比如int, string类型。
*/
func binarySize(i int,
	i8 int8,
	i16 int16,
	i32 int32,
	i64 int64,
	f32 float32,
	f64 float64,
	b bool,
	s string,
	listI64 []int64,
	pi *int) {

	// int类型的字节数跟操作系统相关
	// 在8位系统下是1， 在16位系统下2， 在32位系统下4， 在64位系统下是8
	// 所以这里使用binary.Size计算的长度是-1
	fmt.Printf("%T size: %d\n", i, binary.Size(i))

	fmt.Printf("%T size: %d\n", i8, binary.Size(i8))   // 1
	fmt.Printf("%T size: %d\n", i16, binary.Size(i16)) // 2
	fmt.Printf("%T size: %d\n", i32, binary.Size(i32)) // 4
	fmt.Printf("%T size: %d\n", i64, binary.Size(i64)) // 8
	fmt.Printf("%T size: %d\n", f32, binary.Size(f32)) // 4
	fmt.Printf("%T size: %d\n", f64, binary.Size(f64)) // 8

	// 布尔类型在任何系统下都是用一个字节表示
	fmt.Printf("%T size: %d\n", b, binary.Size(b))

	// string类型不是固定长度的数据类型，所以计算长度是-1
	fmt.Printf("%T size: %d\n", s, binary.Size(s))

	// [1 2] size: 16
	// 计算的是实际包含的两个int64数值的字节数，这个结果是正确的
	fmt.Printf("%v size: %d\n", listI64, binary.Size(listI64))

	// 指针类型的字节数跟操作系统相关
	// 在8位系统下是1， 在16位系统下2， 在32位系统下4， 在64位系统下是8
	// 所以这里使用binary.Size计算的长度是-1
	fmt.Printf("%T size: %d\n", pi, binary.Size(pi))
}

func unsafeSizeof(i int,
	i8 int8,
	i16 int16,
	i32 int32,
	i64 int64,
	f32 float32,
	f64 float64,
	b bool,
	s string,
	listI64 []int64,
	pi *int) {

	fmt.Printf("%T size: %d\n", i, unsafe.Sizeof(i))
	fmt.Printf("%T size: %d\n", i8, unsafe.Sizeof(i8))
	fmt.Printf("%T size: %d\n", i16, unsafe.Sizeof(i16))
	fmt.Printf("%T size: %d\n", i32, unsafe.Sizeof(i32))
	fmt.Printf("%T size: %d\n", i64, unsafe.Sizeof(i64))
	fmt.Printf("%T size: %d\n", f32, unsafe.Sizeof(f32))
	fmt.Printf("%T size: %d\n", f64, unsafe.Sizeof(f64))
	fmt.Printf("%T size: %d\n", b, unsafe.Sizeof(b))

	// string size: 16
	// 为什么是16？
	// 查看Go语言的源码src/runtime/string.go，可以发现string类型的数据结构如下：
	/*
		type stringStruct struct {
			str unsafe.Pointer
			len int
		}
	*/
	// 结构中包含两个字段：一个表示首地址的指针类型，一个表示字符串长度的int类型
	// 在我的64位机器上字节数刚好是16
	fmt.Printf("%T size: %d\n", s, unsafe.Sizeof(s))

	// [1 2] size: 24
	// 为什么是24？
	// 查看Go语言的源码src/runtime/slice.go，可以发现slice类型的数据结构如下：
	/*
		type slice struct {
			array unsafe.Pointer
			len   int
			cap   int
		}
	*/
	// 结构中包含三个字段：一个表示首地址的指针，一个表示长度的int类型，一个表示容量的int类型
	// 在我的64位机器上字节数刚好是24
	fmt.Printf("%v size: %d\n", listI64, unsafe.Sizeof(listI64))

	fmt.Printf("%T size: %d\n", pi, unsafe.Sizeof(pi))
}

func main() {
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	var f32 float32
	var f64 float64
	var b bool
	var s = "hello"
	var listI64 []int64
	listI64 = append(listI64, 1)
	listI64 = append(listI64, 2)
	var pi = &i

	// Use binary.Size to calculate the variable size
	fmt.Println("Use binary.Size to calculate the variable size")
	binarySize(i, i8, i16, i32, i64, f32, f64, b, s, listI64, pi)
	fmt.Println()

	// Use unsafe.Sizeof to calculate the variable size
	fmt.Println("Use unsafe.Sizeof to calculate the variable size")
	unsafeSizeof(i, i8, i16, i32, i64, f32, f64, b, s, listI64, pi)
	fmt.Println()
}
