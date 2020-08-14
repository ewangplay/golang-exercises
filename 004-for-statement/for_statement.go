package main

import (
	"fmt"
)

func main() {
	//sum 1 to 10
	fmt.Println("sum 1 to 10: ", sum(10))

	//revert string hello
	fmt.Println("reverse of hello: ", revert("hello"))

	//range test
	list := []string{"a", "b", "c", "d"}
	for k, v := range list {
		fmt.Println("key = ", k, "value = ", v)
	}

    s := "hello"
	for i, c := range s {
		fmt.Printf("char %c at %d index\n", c, i)
	}
}

func sum(num int) int {
	sum := 0
	for i := 0; i < num; i++ {
		sum += i
	}
	return sum
}

func revert(s string) string {
	bstr := []byte(s)
	for i, j := 0, len(bstr)-1; i < j; i, j = i+1, j-1 {
		//为什么可以这样使用呢？看来平行赋值没有先后赋值顺序，而是同时进行的
		//也是说明golang中的确没有逗号表达式，逗号表达式中是有前后顺序的，平行
		//赋值只是golang中一个特殊的赋值方式
		//但平行赋值在这里是怎么做到同时赋值的，中间有没有使用临时变量，还有待于
		//进一步考察。
		bstr[i], bstr[j] = bstr[j], bstr[i]
	}
	return string(bstr)
}
