// 单例模式

package main

import (
	"fmt"
	"sync"
)

type Singleton struct{}

var instance *Singleton
var once sync.Once

func NewSingleton() *Singleton {
	// 只会执行一次
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func main() {
	var obj2 *Singleton

	wait := make(chan int)
	obj := NewSingleton()

	// 启动一个新的 goroutine 调用 NewSingleton 函数创建 Singleton 的实例
	go func() {
		defer func() {
			wait <- 0
		}()
		obj2 = NewSingleton()
	}()
	<-wait
	close(wait)

	fmt.Println(obj == obj2)
}
