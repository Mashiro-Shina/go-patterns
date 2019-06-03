// 命令模式

package main

import "fmt"

// 统一命令接口
type Command interface {
	execute()
}

// 具体的命令
type ConcretCommand1 struct {
	receiver *Receiver
}

// 命令执行方法
func (c *ConcretCommand1) execute() {
	c.receiver.do1()
}

type ConcretCommand2 struct {
	receiver *Receiver
}

func (c *ConcretCommand2) execute() {
	c.receiver.do2()
}

// 命令的实际执行者
type Receiver struct{}

func (receiver *Receiver) do1() {
	fmt.Println("doing concret command1")
}

func (receiver *Receiver) do2() {
	fmt.Println("doing concret command2")
}

// 命令的请求者
type Invoker struct {
	command Command
}

// 设置命令
func (invoker *Invoker) setCommand(command Command) {
	invoker.command = command
}

// 执行命令
func (invoker *Invoker) runCommand() {
	invoker.command.execute()
}

func main() {
	invoker := &Invoker{}
	receiver := &Receiver{}

	invoker.setCommand(&ConcretCommand1{receiver})
	invoker.runCommand()

	invoker.setCommand(&ConcretCommand2{receiver})
	invoker.runCommand()
}
