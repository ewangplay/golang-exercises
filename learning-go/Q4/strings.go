/*
<<Learning Go>>
Q4: Strings
1. Create a Go program that prints the following (up to 100 characters):
A
AA
AAA
AAAA
AAAAA
AAAAAA
AAAAAAA
...

2. Create a Go program that counts the number of characters in this string:
asSASA ddd dsjkdsjs dk

3. Extend the program from previous question to replace the three runes at position 4 with "abc"

4. Write a Go program that reverses a string, so “foobar” is printed as “raboof”.
*/
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//Q4-1
	func1()

	//Q4-2
	func2_1()
	func2_2()

	//Q4-3
	func3()

	//Q4-4
	func4()
}

func func1() {
	str := "A"
	for i := 1; i <= 100; i++ {
		fmt.Println(str)
		str += "A"
	}
	fmt.Println()
}

func func2_1() {
	s := "asSASA ddd dsjkdsjs dk"
	m := make(map[string]int)

	for _, v := range s {
		m[string(v)]++
	}

	for k, v := range m {
		fmt.Printf("character %s num: %d\n", k, v)
	}
	fmt.Println()
}

func func2_2() {
	s := "asSASA ddd dsjkdsjs dk 我的中国心"
	fmt.Printf("String [%v] Len: %v, Characters: %v\n", s, len([]byte(s)), utf8.RuneCount([]byte(s)))
	fmt.Println()
}

func func3() {
	s := "asSASA ddd dsjkdsjs dk"
	s1 := []byte(s)
	n1 := copy(s1[4:], []byte("abc"))
	fmt.Printf("replace characters num %d, new string is %s\n", n1, s1)
	fmt.Println()
}

func func4() {
	s := "foobar"
	a := []byte(s)

	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	fmt.Println(string(a))
}
