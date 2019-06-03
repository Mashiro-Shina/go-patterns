// 享元模式

package main

import "fmt"

// 通用的享元接口
type FlyWeight interface {
	process()
}

// 一个具体的享元结构
type ConcretFlyWeight struct {
	name string
}

func (flyweight *ConcretFlyWeight) process() {
	fmt.Printf("ConcretFlyWeight %s\n", flyweight.name)
}

// 一个享元工厂，用来创建并管理享元对象
type FlyWeightFactory struct {
	flyWeights map[string]FlyWeight
}

func (factory *FlyWeightFactory) newFlyWeight(name string) FlyWeight {
	if instance, ok := factory.flyWeights[name]; ok {
		fmt.Printf("Getting FlyWeight %s from the pool\n", name)
		return instance
	} else {
		fmt.Printf("Creating FlyWeight %s\n", name)
		factory.flyWeights[name] = &ConcretFlyWeight{name}
		return factory.flyWeights[name]
	}
}

func main() {
	factory := FlyWeightFactory{make(map[string]FlyWeight)}
	factory.newFlyWeight("instance1").process()
	factory.newFlyWeight("instance1").process()
}
