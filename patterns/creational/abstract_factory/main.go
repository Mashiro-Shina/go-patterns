// 抽象工厂模式

package main

import "fmt"

// Person 接口
type Person interface {
	say()
}

// 学生结构
type Student struct {
	name  string
	class string
}

func (student *Student) say() {
	fmt.Printf("I'm %s, Grade %s\n", student.name, student.class)
}

// 教师结构
type Teacher struct {
	name    string
	subject string
}

func (teacher *Teacher) say() {
	fmt.Printf("I'm %s, I teach %s\n", teacher.name, teacher.subject)
}

// 医生结构
type Doctor struct {
	name     string
	hospital string
}

func (doctor *Doctor) say() {
	fmt.Printf("I'm %s, I work in %s\n", doctor.name, doctor.hospital)
}

// 用于创建实现 Person 接口的结构体实例的工厂
type PersonFactory struct{}

func (factory *PersonFactory) NewPerson(target string, attributions ...string) Person {
	var instance Person

	switch target {
	case "student":
		instance = &Student{
			name:  attributions[0],
			class: attributions[1],
		}
	case "teacher":
		instance = &Teacher{
			name:    attributions[0],
			subject: attributions[1],
		}
	case "doctor":
		instance = &Doctor{
			name:     attributions[0],
			hospital: attributions[1],
		}
	default:
		fmt.Printf("can not create instance for target %s\n", target)
	}

	return instance
}

// 动物接口
type Animal interface {
	run()
}

// 狗结构体
type Dog struct {
	name string
}

func (dog *Dog) run() {
	fmt.Printf("Dog %s is running\n", dog.name)
}

// 猫结构体
type Cat struct {
	name string
}

func (cat *Cat) run() {
	fmt.Printf("Cat %s is running\n", cat.name)
}

// 老虎结构体
type Tiger struct {
	name string
}

func (tiger *Tiger) run() {
	fmt.Printf("Tiger %s is running\n", tiger.name)
}

// 动物工厂
type AnimalFactory struct{}

func (factory *AnimalFactory) NewAnimal(target string, attributions ...string) Animal {
	var instance Animal

	switch target {
	case "dog":
		instance = &Dog{
			name: attributions[0],
		}
	case "cat":
		instance = &Cat{
			name: attributions[0],
		}
	case "tiger":
		instance = &Tiger{
			name: attributions[0],
		}
	default:
		fmt.Printf("can not create instance for target %s\n", target)
	}

	return instance
}

// 抽象工厂
type AbstractFactory struct{}

func (factory *AbstractFactory) NewFactory(target string) interface{} {
	var instance interface{}

	switch target {
	case "person":
		instance = &PersonFactory{}
	case "animal":
		instance = &AnimalFactory{}
	default:
		fmt.Printf("can not create instance for %s\n", target)
	}
	return instance
}

func main() {
	factory := &AbstractFactory{}

	personFactory := factory.NewFactory("person").(*PersonFactory)
	personFactory.NewPerson("student", "jack", "2017").say()
	personFactory.NewPerson("teacher", "jack", "English").say()
	personFactory.NewPerson("doctor", "jack", "People's Hospital").say()

	animalFactory := factory.NewFactory("animal").(*AnimalFactory)
	animalFactory.NewAnimal("dog", "dog").run()
	animalFactory.NewAnimal("cat", "cat").run()
	animalFactory.NewAnimal("tiger", "tiger").run()
}
