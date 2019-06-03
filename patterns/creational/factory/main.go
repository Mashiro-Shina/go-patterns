// 工厂模式

package main

import "fmt"

// 通用接口
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
	fmt.Printf("I'm %s, I work in %s", doctor.name, doctor.hospital)
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

func main() {
	factory := PersonFactory{}
	factory.NewPerson("student", "jack", "2017").say()
	factory.NewPerson("teacher", "jack", "English").say()
	factory.NewPerson("doctor", "jack", "People's Hospital").say()
}
