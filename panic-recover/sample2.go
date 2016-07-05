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

	r, err := Int8FromInt64(n)
	if err != nil {
		fmt.Printf("转换%v到类型int8失败：%v\n", n, err)
		os.Exit(1)
	}
	fmt.Printf("转换%v到类型int8成功: %v\n", n, r)
}

func Int8FromInt64(num int64) (r int8, err error) {
	//如果有多个defer调用，那么这些defer调用的执行策略是LIFO
	defer func() {
		if err != nil {
            err = fmt.Errorf("这是一个修改后error: %v", err)
		}
	}()

	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()

	r = ConvertInt64ToInt8(num)

	return r, nil
}

func ConvertInt64ToInt8(num int64) int8 {
	if math.MinInt8 <= num && num <= math.MaxInt8 {
		return int8(num)
	}
	panic("out of the int8 range")
}
