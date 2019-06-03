// 状态模式

package main

import "fmt"

// 统一状态接口
type State interface {
	do()
}

// 具体的状态
type Eating struct {
	meal string
}

func (e *Eating) do() {
	fmt.Printf("Eating %s\n", e.meal)
}

type Drinking struct {
	drinks string
}

func (d *Drinking) do() {
	fmt.Printf("Drinking %s\n", d.drinks)
}

type Sleeping struct {
	bedroom string
}

func (s *Sleeping) do() {
	fmt.Printf("Sleeping in %s\n", s.bedroom)
}

// context 对象
type Person struct {
	state    State
	eating   *Eating
	drinking *Drinking
	sleeping *Sleeping
}

// 创建一个新的 Person 实例
func NewPerson(eating *Eating, drinking *Drinking, sleeping *Sleeping) *Person {
	return &Person{
		state:    eating,
		eating:   eating,
		drinking: drinking,
		sleeping: sleeping,
	}
}

// 改变 person 状态
func (person *Person) switchTo(target string) {
	switch target {
	case "eating":
		person.state = person.eating
	case "drinking":
		person.state = person.drinking
	case "sleeping":
		person.state = person.sleeping
	default:
		fmt.Printf("No such target state: \"%s\"", target)
	}
}

func (person *Person) do() {
	person.state.do()
}

func main() {
	eating := &Eating{"rice"}
	drinking := &Drinking{"water"}
	sleeping := &Sleeping{"my bedroom"}
	person := NewPerson(eating, drinking, sleeping)

	person.do()

	for _, state := range []string{"drinking", "sleeping", "eating"} {
		person.switchTo(state)
		person.do()
	}
}
