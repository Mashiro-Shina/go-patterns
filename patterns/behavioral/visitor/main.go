// 访问者模式

package main

import "fmt"

// 数字节点
type Number struct {
	value float64
}

// 相反数
type Negate struct {
	operand interface{}
}

// 加法
type Add struct {
	left, right interface{}
}

// 减法
type Sub struct {
	left, right interface{}
}

// 乘法
type Mul struct {
	left, right interface{}
}

// 除法
type Div struct {
	left, right interface{}
}

// 访问者结构
type Visitor struct{}

// 访问方法
func (visitor *Visitor) visit(target interface{}) float64 {
	// 根据访问的目标类型处理
	switch target.(type) {
	case Number:
		// Number 直接返回其 value
		number := target.(Number)
		return number.value
	case Negate:
		// 对递归后的结果取相反数
		negate := target.(Negate)
		return -visitor.visit(negate.operand)
	case Add:
		// 对左右操作数分别递归然后相加
		add := target.(Add)
		return visitor.visit(add.left) + visitor.visit(add.right)
	case Sub:
		sub := target.(Sub)
		return visitor.visit(sub.left) - visitor.visit(sub.right)
	case Mul:
		mul := target.(Mul)
		return visitor.visit(mul.left) * visitor.visit(mul.right)
	case Div:
		div := target.(Div)
		return visitor.visit(div.left) / visitor.visit(div.right)
	default:
		fmt.Println("can not process the target.")
		return -1
	}
}

func main() {
	visitor := Visitor{}

	t1 := Div{Number{6}, Number{3}}
	t2 := Sub{Number{5}, t1}
	t3 := Mul{t2, Number{8}}
	t4 := Add{Number{10}, t3}
	t5 := Add{Negate{Number{1}}, t4}

	fmt.Println(visitor.visit(t5))
}
