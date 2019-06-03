// 根据传入的参数决定调用的函数

package main

import "fmt"

type Catalog struct {
	methods map[string]func() // 使用一个映射来根据参数选择函数
}

func NewCatalog() *Catalog {
	catalog := &Catalog{}

	// 初始化 methods 映射
	catalog.methods = map[string]func(){
		"arg1": catalog.func1,
		"arg2": catalog.func2,
	}
	return catalog
}

func (catalog *Catalog) run(arg string) {
	// 执行入口

	if method, ok := catalog.methods[arg]; ok {
		method()
	} else {
		fmt.Printf("no such method: <%s>", arg)
	}
}

func (catalog *Catalog) func1() {
	fmt.Println("function 1")
}

func (catalog *Catalog) func2() {
	fmt.Println("function 2")
}

func main() {
	catalog := NewCatalog()
	catalog.run("arg1")
	catalog.run("arg2")
	catalog.run("not exists")
}
