/* 这个例子演示了在Go语言中如何实现继承关系 */

package main

import (
	"fmt"
)

type Animal struct {
	Name string
}

/* 这里定义了一个动物通用的Eat方法 */
func (this *Animal) Eat() {
	fmt.Printf("I am an animal. My name is %s. I want to eat something.\n", this.Name)
}

/* Dog类型通过嵌入Animal类型继承了Animal的Eat行为 */
type Dog struct {
	Animal
}

/* Cat类型通过嵌入Animal类型继承了Animal的Eat行为 */
type Cat struct {
	Animal
}

/* 但是Cat类型重写(override)了Eat行为 */
func (this *Cat) Eat() {
	fmt.Printf("I am a cat. My name is %s. I want to eat fish.\n", this.Name)
}

func main() {
	itsAnimal := Animal{Name: "Gaga"}
	itsAnimal.Eat()

	itsDog := Dog{Animal: Animal{Name: "Waly"}}
	itsDog.Eat()

	itsCat := Cat{Animal: Animal{Name: "Solory"}}
	itsCat.Eat()
}
