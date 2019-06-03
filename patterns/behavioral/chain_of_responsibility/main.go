// 责任链模式

package main

import "fmt"

type Handler interface {
	check(request int) bool
	handle(request int)
}

type Handler0 struct {
	successor Handler // 责任链中的下一个处理器，是一个 Handler 接口类型
}

func (handler *Handler0) check(request int) bool {
	// 判断该处理器是否可以处理请求
	if request >= 0 && request < 10 {
		return true
	}
	return false
}

func (handler *Handler0) handle(request int) {
	if handler.check(request) {
		// 进行详细的处理
		fmt.Printf("request %d handled in handler 0\n", request)
	} else {
		// 将请求传递给责任链中下一个处理器
		handler.successor.handle(request)
	}
}

type Handler1 struct {
	successor Handler
}

func (handler *Handler1) check(request int) bool {
	if request >= 10 && request < 20 {
		return true
	}
	return false
}

func (handler *Handler1) handle(request int) {
	if handler.check(request) {
		fmt.Printf("request %d handled in handler 1\n", request)
	} else {
		handler.successor.handle(request)
	}
}

type FallbackHandler struct{}

func (handler *FallbackHandler) check(request int) bool {
	return false
}

func (handler *FallbackHandler) handle(request int) {
	fmt.Printf("end of chain, no handler for %d\n", request)
}

func main() {
	handler0 := &Handler0{}
	handler1 := &Handler1{}
	fallback := &FallbackHandler{}
	handler0.successor = handler1
	handler1.successor = fallback

	// 所有请求
	requests := []int{3, 9, 10, 18, 22, 1}
	for _, request := range requests {
		// 调用第一个处理器处理
		handler0.handle(request)
	}
}
