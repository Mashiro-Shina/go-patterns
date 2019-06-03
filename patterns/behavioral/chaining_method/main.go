package main

import "fmt"

type Person struct {
	name, gender string
	age          int
	action       *Action
}

func (person *Person) showInfo() *Action {
	fmt.Printf("Person(%s, %s, %d)\n", person.name, person.gender, person.age)
	return person.action
}

type Action struct {
	name string
}

func (action *Action) run() *Action {
	fmt.Printf("doing action %s\n", action.name)
	return action
}

func (action *Action) stop() {
	fmt.Printf("stopping action %s\n", action.name)
}

func main() {
	person := &Person{
		name:   "Jack",
		gender: "male",
		age:    18,
		action: &Action{
			name: "move",
		},
	}
	person.showInfo().run().stop()
}
