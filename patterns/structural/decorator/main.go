// 装饰器模式

package main

import "fmt"

type Shape interface {
	draw()
}

type Rectangle struct {
	width  float64
	height float64
}

func (rectangle *Rectangle) draw() {
	fmt.Printf("Drawing shape: Rectangle<width=%.2f, height=%.2f>\n", rectangle.width, rectangle.height)
}

func (rectangle *Rectangle) String() string {
	return fmt.Sprintf("Rectangle<width=%.2f, height=%.2f>", rectangle.width, rectangle.height)
}

type Circle struct {
	x, y   float64
	radius float64
}

func (circle *Circle) draw() {
	fmt.Printf("Drawing shape: Circle<x=%.2f, y=%.2f, radius=%.2f>\n", circle.x, circle.y, circle.radius)
}

func (circle *Circle) String() string {
	return fmt.Sprintf("Circle<x=%.2f, y=%.2f, radius=%.2f>", circle.x, circle.y, circle.radius)
}

// 一个装饰器结构
type ColorDecorator struct {
	shape Shape
}

func (decorator *ColorDecorator) draw(color string) {
	decorator.shape.draw()

	// 为 Shape 添加颜色
	decorator.setColor(decorator.shape, color)
}

func (decorator *ColorDecorator) setColor(shape Shape, color string) {
	fmt.Printf("set %s's color: %s\n", shape, color)
}

func main() {
	rectangle := &Rectangle{1.1, 2.2}
	coloredRectangle := &ColorDecorator{rectangle}
	coloredRectangle.draw("red")

	circle := &Circle{1.1, 2.2, 3.3}
	coloredCircle := &ColorDecorator{circle}
	coloredCircle.draw("blue")
}
