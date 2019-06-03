// 外观模式

package main

import (
	"fmt"
	"time"
)

// 引擎结构体
type Engine struct{}

func (engine *Engine) start() {
	fmt.Println("Engine started")
}

func (engine *Engine) stop() {
	fmt.Println("Engine stopped")
}

// 档位结构体
type Gear struct {
	maxGear     int
	currentGear int
}

func (gear *Gear) shiftUp() {
	if gear.currentGear == gear.maxGear {
		fmt.Println("Can not shift up anymore")
	} else {
		gear.currentGear++
		fmt.Printf("Shifted up to gear %d\n", gear.currentGear)
	}
}

func (gear *Gear) shiftDown() {
	if gear.currentGear == -1 {
		fmt.Println("In reverse, can not shift down")
	} else {
		gear.currentGear--
		fmt.Printf("Shifted down to gear %d\n", gear.currentGear)
	}
}

func (gear *Gear) shiftReverse() {
	gear.currentGear = -1
	fmt.Println("reverse shifting")
}

func (gear *Gear) shiftTo(target int) {
	gear.currentGear = target
	fmt.Printf("Shifted to gear %d\n", gear.currentGear)
}

// 手刹结构体
type Brake struct{}

func (brake *Brake) engage() {
	fmt.Println("Brake engaged")
}

func (brake *Brake) release() {
	fmt.Println("Brake released")
}

// 汽车结构体
type Car struct {
	engine *Engine
	gear   *Gear
	brake  *Brake
}

// 启动汽车
func (car *Car) start() {
	fmt.Println("Starting the car")
	car.brake.release()
	car.engine.start()
	car.gear.shiftUp()
	fmt.Println("Car started")
	car.gear.shiftTo(3)
}

// 停止汽车
func (car *Car) stop() {
	fmt.Println("Stopping the car")
	car.gear.shiftTo(0)
	car.engine.stop()
	car.brake.engage()
	fmt.Println("Car stopped")
}

func main() {
	car := &Car{
		engine: &Engine{},
		gear:   &Gear{5, 0},
		brake:  &Brake{},
	}

	car.start()
	time.Sleep(3 * time.Second)
	car.stop()
}
