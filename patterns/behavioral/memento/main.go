// 备忘录模式

package main

import "fmt"

func memento(target *Number) *func() {
	// 备忘录函数

	// 创建一个新的对象，它的值和传入的对象一致
	state := new(Number)
	*state = *target

	// 调用这个函数时恢复对象的状态
	refresh := func() {
		*target = *state
	}

	return &refresh
}

type Transation struct {
	targets []*Number // 一个 Number 指针切片
	states  []*func() // 一个状态切片，元素类型为函数指针
}

func NewTransaction(objs ...*Number) *Transation {
	// 构造一个新的 Transaction 实例，并返回一个指针

	targets := []*Number{}
	for _, obj := range objs {
		targets = append(targets, obj)
	}
	return &Transation{
		targets: targets,
	}
}

func (transaction *Transation) commit() {
	// 保存 targets 属性中所有对象的状态

	for _, target := range transaction.targets {
		transaction.states = append(transaction.states, memento(target))
	}
}

func (transaction *Transation) rollback() {
	// 将所有对象的状态重置

	for _, state := range transaction.states {
		(*state)()
	}
}

type Number struct {
	value int
}

func (number *Number) increment() {
	number.value++
}

func (number *Number) String() string {
	return fmt.Sprintf("Number(%d)", number.value)
}

func main() {
	var number *Number
	var transaction *Transation

	number = &Number{-1}
	transaction = NewTransaction(number)

	fmt.Println(number)

	// 提交所有状态
	transaction.commit()

	for i := 0; i < 3; i++ {
		number.increment()
	}
	fmt.Println(number)

	// 所有对象状态都回滚到提交的状态
	transaction.rollback()
	fmt.Println(number)
}
