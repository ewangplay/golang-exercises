/*
这个例子演示了在Go语言中如何实现继承关系
*/
package main

import (
	"fmt"
)

// Animal 定义了一个"动物"类
type Animal struct {
	name string
}

// SetName 设定动物的名字
func (a *Animal) SetName(name string) {
	a.name = name
}

// GetName 获取动物的名字
func (a *Animal) GetName() string {
	return a.name
}

// Eat 定义动物通用的"吃"方法
func (a *Animal) Eat() {
	fmt.Printf("I am an animal. My name is %s. I want to eat something.\n", a.GetName())
}

// Dog 定义了一个"狗"类，通过嵌入Animal类型继承了Animal的行为
type Dog struct {
	Animal
}

// Cat 定义了一个"猫"类，通过嵌入Animal类型继承了Animal的行为
type Cat struct {
	Animal
}

// Eat 重写(override)了Eat行为
func (c *Cat) Eat() {
	fmt.Printf("I am a cat. My name is %s. I want to eat fish.\n", c.GetName())
}

func main() {
	a := Animal{}
	a.SetName("Gaga")
	a.Eat()

	d := Dog{}
	d.SetName("Waly")
	d.Eat()

	c := Cat{}
	c.SetName("Solory")
	c.Eat()
}
