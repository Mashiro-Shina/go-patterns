// 原型模式

package main

import "fmt"

type Prototype interface {
	clone() Prototype
}

type Person struct {
	name string
}

func (person *Person) say() {
	fmt.Printf("Hi, I'm %s", person.name)
}

// 克隆方法
func (person *Person) clone() *Person {
	newPerson := new(Person)
	*newPerson = *person
	return newPerson
}

func main() {
	person := &Person{"Jack"}
	person2 := person.clone()

	fmt.Println(*person == *person2)
}
