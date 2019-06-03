// 桥接模式

package main

import "fmt"

// 桥接实现接口
type DrawingAPI interface {
	drawCircle(x, y, radius float64)
}

// 具体实现结构
type ConcretDrawingAPI1 struct{}

func (api *ConcretDrawingAPI1) drawCircle(x, y, radius float64) {
	fmt.Printf("DrawingAPI1.Circle at %.2f:%.2f, radius: %.2f\n", x, y, radius)
}

type ConcretDrawingAPI2 struct{}

func (api *ConcretDrawingAPI2) drawCircle(x, y, radius float64) {
	fmt.Printf("DrawingAPI2.Circle at %.2f:%.2f, radius: %.2f\n", x, y, radius)
}

type CircleShape struct {
	x, y       float64
	radius     float64
	drawingAPI DrawingAPI
}

func NewCircleShape(x, y, radius float64, api DrawingAPI) *CircleShape {
	return &CircleShape{
		x:          x,
		y:          y,
		radius:     radius,
		drawingAPI: api,
	}
}

func (circle *CircleShape) draw() {
	circle.drawingAPI.drawCircle(circle.x, circle.y, circle.radius)
}

func main() {
	circle := NewCircleShape(4.0, 5.0, 3.0, &ConcretDrawingAPI1{})
	circle.draw()

	circle = NewCircleShape(2.0, 3.0, 4.0, &ConcretDrawingAPI2{})
	circle.draw()
}
