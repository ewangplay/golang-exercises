/*
在多级函数调用的场景下，底层的一个函数抛出一个异常，这个异常会逐级向上层传递。

如果上层调用者在defer中使用recover捕获这个异常，就可以对这个异常进行处理；
如果上层调用者没有在defer中使用recover捕获这个异常，那么异常继续向更上层调用者传递，
以此类推，直到到达main函数。

如果main函数在defer中使用recover捕获了这个异常，就可以对这个异常进行处理；
如果main函数没有在defer中使用recover捕获这个异常，那么这个异常就到此终止，
程序异常终止，panic的信息会在dump的程序调用栈中显示出来。
*/
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
			fmt.Printf("这里是异常的终点站: %v\n程序将退出...\n", e)
			os.Exit(1)
		}
	}()

	r, err := int8FromInt64(n)
	if err != nil {
		fmt.Printf("转换%v到类型int8失败：%v\n", n, err)
		os.Exit(1)
	}
	fmt.Printf("转换%v到类型int8成功: %v\n", n, r)
}

func int8FromInt64(num int64) (r int8, err error) {
	// 如果一个函数中有多个defer调用，这些defer调用顺序策略是LIFO(后进先出)
	// 返回的error在这里被截获，修改后作为异常被重新抛出
	defer func() {
		if err != nil {
			err = fmt.Errorf("要想从此过，留下买路财： %v", err)
			panic(err)
		}
	}()

	// 底层抛出的异常首先在这里被捕获，被转换成了error返回
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("int8FromInt64 - 我捕获了一个异常: %v", e)
		}
	}()

	r = convertInt64ToInt8(num)

	return r, nil
}

func convertInt64ToInt8(num int64) int8 {
	if math.MinInt8 <= num && num <= math.MaxInt8 {
		return int8(num)
	}
	// 这里抛出了一个异常
	panic("convertInt64ToInt8 - out of the int8 range")
}
