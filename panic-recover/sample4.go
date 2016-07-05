package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %v <number>\n", os.Args[0])
		os.Exit(1)
	}

	strN := os.Args[1]
	n, err := strconv.ParseInt(strN, 0, 0)
	if err != nil {
		fmt.Printf("input number invalid. %v\n", err)
		os.Exit(1)
	}

	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("捕获到异常: %v, 程序将退出.\n", e)
			os.Exit(1)
		}
	}()

	r, err := Int8FromInt64(n)
	if err != nil {
		fmt.Printf("转换%v到类型int8失败：%v\n", n, err)
		os.Exit(1)
	}
	fmt.Printf("转换%v到类型int8成功: %v\n", n, r)
}

func Int8FromInt64(num int64) (r int8, err error) {
	//在多级函数调用的场景下，底层的一个函数调用了panic抛出一个异常，
	//这个异常就会向上层调用者传递，如果上层调用者在defer调用中使用了recover捕获这个异常，就可以对这个异常进行处理；
	//如果上层调用者没有在defer中使用recover捕获这个异常，那么这个异常不会到此终止，而是继续向更上层调用者传递，以此类推。
	//直到到达main函数，如果main函数中在defer中使用recover捕获了这个异常，就可以对这个异常进行处理；如果main函数没有在
	//defer中使用recover捕获这个异常，那么这个异常就到此终止，程序也异常终止，带有最初的传给panic的信息会在dump的程序
	//调用栈中显示出来，以提醒开发者注意。
	/*
		defer func() {
			if e := recover(); e != nil {
				panic(fmt.Sprintf("Int8FromInt64: %v", e))
			}
		}()
	*/

	r = ConvertInt64ToInt8(num)

	return r, nil
}

func ConvertInt64ToInt8(num int64) int8 {
	if math.MinInt8 <= num && num <= math.MaxInt8 {
		return int8(num)
	}
	panic("out of the int8 range")
}
