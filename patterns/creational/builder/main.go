// 建造者模式

package main

import "fmt"

type House struct {
	n_windows  int        // 窗户的数目
	n_doors    int        // 门的数目
	n_bedrooms int        // 卧室的数目
	windows    []*Window  // 所有窗户
	doors      []*Door    // 所有门
	bedrooms   []*Bedroom // 所有卧室
}

func NewHouse(n_windows, n_doors, n_bedrooms int) *House {
	// 新建一个 House 对象，并返回一个指针

	return &House{
		n_windows:  n_windows,
		n_doors:    n_doors,
		n_bedrooms: n_bedrooms,
	}
}

func (house *House) AddWindow(window *Window) {
	// 向 house 添加一个 window 指针

	house.windows = append(house.windows, window)
}

func (house *House) AddDoor(door *Door) {
	// 向 house 添加一个 door 指针

	house.doors = append(house.doors, door)
}

func (house *House) AddBedroom(bedroom *Bedroom) {
	// 向 house 添加一个 bedroom 指针

	house.bedrooms = append(house.bedrooms, bedroom)
}

func (house *House) String() string {
	// 打印时返回字符串

	result := fmt.Sprintf("House(windows=%d, doors=%d, bedrooms=%d)\n", house.n_windows, house.n_doors, house.n_bedrooms)
	for _, window := range house.windows {
		result += fmt.Sprintf("\t%s\n", window)
	}
	for _, door := range house.doors {
		result += fmt.Sprintf("\t%s\n", door)
	}
	for _, bedroom := range house.bedrooms {
		result += fmt.Sprintf("\t%s\n", bedroom)
	}
	return result
}

type Window struct {
	width  float64
	height float64
}

func NewWindow(width, height float64) *Window {
	return &Window{
		width:  width,
		height: height,
	}
}

func (window *Window) String() string {
	return fmt.Sprintf("Window<width=%.2f, height=%.2f>", window.width, window.height)
}

type Door struct {
	width  float64
	height float64
}

func NewDoor(width, height float64) *Door {
	return &Door{
		width:  width,
		height: height,
	}
}

func (door *Door) String() string {
	return fmt.Sprintf("Door<width=%.2f, height=%.2f>", door.width, door.height)
}

type Bedroom struct {
	width  float64
	height float64
}

func NewBedroom(width, height float64) *Bedroom {
	return &Bedroom{
		width:  width,
		height: height,
	}
}

func (bedroom *Bedroom) String() string {
	return fmt.Sprintf("Bedroom<windth=%.2f, height=%.2f>", bedroom.width, bedroom.height)
}

func main() {
	house := NewHouse(2, 1, 3)
	window1 := NewWindow(1.1, 2.2)
	window2 := NewWindow(2.2, 3.3)
	door := NewDoor(1.1, 2.2)
	bedroom1 := NewBedroom(1.1, 2.2)
	bedroom2 := NewBedroom(2.2, 3.3)
	bedroom3 := NewBedroom(3.3, 4.4)
	house.AddWindow(window1)
	house.AddWindow(window2)
	house.AddDoor(door)
	house.AddBedroom(bedroom1)
	house.AddBedroom(bedroom2)
	house.AddBedroom(bedroom3)

	fmt.Println(house)
}
