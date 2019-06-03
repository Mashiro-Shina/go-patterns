// 过滤器模式

package main

import "fmt"

// 需要过滤的结构
type Person struct {
	name   string
	gender string
	age    int
}

// 统一过滤器接口
type Filter interface {
	filter([]Person) []Person
}

// 只选择男性的过滤器
type MenFilter struct{}

func (filter *MenFilter) filter(persons []Person) []Person {
	filteredPersons := []Person{}
	for _, person := range persons {
		if person.gender == "male" {
			filteredPersons = append(filteredPersons, person)
		}
	}

	return filteredPersons
}

// 只选择大于等于 18 岁的人
type AdultFilter struct{}

func (filter *AdultFilter) filter(persons []Person) []Person {
	filteredPersons := []Person{}
	for _, person := range persons {
		if person.age >= 18 {
			filteredPersons = append(filteredPersons, person)
		}
	}

	return filteredPersons
}

func main() {
	persons := []Person{
		Person{"Jack", "male", 18},
		Person{"Amy", "female", 19},
		Person{"Bob", "male", 17},
	}
	menFilter := &MenFilter{}

	men := menFilter.filter(persons)
	fmt.Println(men)

	adultFilter := &AdultFilter{}
	adults := adultFilter.filter(persons)
	fmt.Println(adults)
}
