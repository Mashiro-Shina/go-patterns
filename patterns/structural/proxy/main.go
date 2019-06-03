// 代理模式

package main

import (
	"fmt"
	"sync"
)

// 图片接口
type Image interface {
	display()
}

// 具体的图片结构
type RealImage struct {
	filename string
}

func NewRealImage(filename string) *RealImage {
	realImage := &RealImage{
		filename: filename,
	}
	// 从磁盘中加载图片
	realImage.LoadFromDisk()
	return realImage
}

func (image *RealImage) display() {
	fmt.Printf("Displaying: %s\n", image.filename)
}

func (image *RealImage) LoadFromDisk() {
	fmt.Printf("Loading %s\n", image.filename)
}

type ProxyImage struct {
	once sync.Once

	realImage *RealImage
	filename  string
}

func (image *ProxyImage) display() {
	// 只会执行一次
	image.once.Do(func() {
		image.realImage = NewRealImage(image.filename)
	})

	image.realImage.display()
}

func main() {
	image := &ProxyImage{
		filename: "photo.jpg",
	}
	// 需要加载
	image.display()

	// 不需要加载
	image.display()
}
