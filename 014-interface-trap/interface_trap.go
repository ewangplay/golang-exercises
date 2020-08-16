/*
该实例演示了在使用interface的过程中可能出现的一个很隐蔽的陷阱。
interface类型由两部分构成：dynamic type & dynamic value.
一个interface类型的变量只有当dynamic type 和 dynamic value
两部分同时为nil时，(interface类型的变量 == nil)的结果才为true。
*/
package main

import (
	"bytes"
	"fmt"
	"io"
)

const debug = false

func main() {
	// 修复该问题的方案是定义变量时就定义为interface类型,
	// 而不要定位具体的类型
	// var buf io.Writer // 这是正确的方案
	var buf *bytes.Buffer

	if debug {
		buf = new(bytes.Buffer)
	}

	f(buf) // NOTE: subtly incorrect!

	if debug {
		fmt.Println(buf)
	}
}

// We might expect that changing debug to false
// would disable the collection of the output,
// but in fact it causes the program to panic
// during the out.Write call.
//
// When main calls f, it assigns a nil pointer
// of type *bytes.Buffer to the out parameter,
// so the dynamic value of out is nil.
// However,its dynamic type is *bytes.Buffer,
// meaning that out is a non-nil interface
// containing a nil pointer value, so the
// defensive check out != nil is still true.
func f(out io.Writer) {
	fmt.Printf("%T: %v\n", out, out)
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
