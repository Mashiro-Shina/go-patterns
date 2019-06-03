// 适配器模式

package main

import "fmt"

// SD card 接口
type SDCard interface {
	readSD() string
	writeSD(string)
}

// 一个具体的 SD 卡结构
type ConcretSDCard struct {
	name    string
	content string
}

func (card *ConcretSDCard) readSD() string {
	fmt.Printf("Reading from SD card %s\n", card.name)
	return card.content
}

func (card *ConcretSDCard) writeSD(content string) {
	card.content += content
	fmt.Printf("Write \"%s\" into SD card %s\n", content, card.name)
}

// TF card 接口
type TFCard interface {
	readTF() string
	writeTF(string)
}

// 一个具体的 TF card 结构
type ConcretTFCard struct {
	name    string
	content string
}

func (card *ConcretTFCard) readTF() string {
	fmt.Printf("Reading from TF card %s\n", card.name)
	return card.content
}

func (card *ConcretTFCard) writeTF(content string) {
	card.content += content
	fmt.Printf("Write \"%s\" into TF card %s\n", content, card.name)
}

// computer 接口
type computer interface {
	read()
	write()
}

// 一个具体的电脑结构
type MacBook struct {
	name string
}

// 只能读取 SD 卡的内容
func (macbook *MacBook) read(card SDCard) string {
	return card.readSD()
}

func (macbook *MacBook) write(card SDCard, content string) {
	card.writeSD(content)
}

// 一个适配器结构，用于适配 TFCard，实现 SDCard 接口
type SDAdapterTF struct {
	card TFCard
}

// 读写时调用内部的 TFCard 的读写方法
func (adapter *SDAdapterTF) readSD() string {
	return adapter.card.readTF()
}

func (adapter *SDAdapterTF) writeSD(content string) {
	adapter.card.writeTF(content)
}

func main() {
	macbook := &MacBook{"MacBook"}
	sd := &ConcretSDCard{"SDCard1", "Hello"}

	// 读写 SD 卡
	macbook.read(sd)
	macbook.write(sd, " World")

	// 使用适配器实现读写 TF 卡
	tf := &ConcretTFCard{"TFCard1", "Hello"}
	adapter := &SDAdapterTF{tf}
	macbook.read(adapter)
	macbook.write(adapter, " World")
}
