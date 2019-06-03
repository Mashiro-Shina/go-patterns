// 观察者模式

package main

import "fmt"

// 观察者接口
type Observer interface {
	update(subject *Subject, message string)
}

type Subject struct {
	name      string
	observers []Observer
}

func (subject *Subject) attach(observer Observer) {
	// 订阅
	subject.observers = append(subject.observers, observer)
}

func (subject *Subject) detach(observer Observer) {
	// 取消订阅
	var index int
	for i, obj := range subject.observers {
		if obj == observer {
			index = i
		}
	}
	subject.observers = append(subject.observers[:index], subject.observers[index+1:]...)
}

func (subject *Subject) notify(message string) {
	// 通知所有观察者
	for _, observer := range subject.observers {
		observer.update(subject, message)
	}
}

// 示例观察者
type Person struct {
	name string
}

func (person *Person) update(subject *Subject, message string) {
	fmt.Printf("Person %s: Subject %s got message \"%s\"\n", person.name, subject.name, message)
}

type Dog struct {
	name string
}

func (dog *Dog) update(subject *Subject, message string) {
	fmt.Printf("Dog %s: Subject %s got message \"%s\"\n", subject.name, subject.name, message)
}

func main() {
	subject := &Subject{
		name: "subject",
	}
	person := &Person{"Jack"}
	dog := &Dog{"Husky"}
	subject.attach(person)
	subject.attach(dog)
	subject.notify("Hello")

	subject.detach(dog)
	subject.notify("Hi")
}
