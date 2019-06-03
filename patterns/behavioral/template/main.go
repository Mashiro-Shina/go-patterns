// 模板模式

package main

import "fmt"

// 一个流程接口
type Procedure interface {
	initialize()
	startPlay()
	endPlay()
}

// 一个具体的流程结构
type Football struct{}

func (football *Football) initialize() {
	fmt.Println("Football Game Initialized")
}

func (football *Football) startPlay() {
	fmt.Println("Football Game Started")
}

func (football *Football) endPlay() {
	fmt.Println("Football Game Finished")
}

type Basketball struct{}

func (basketball *Basketball) initialize() {
	fmt.Println("Basketball Game Initialized")
}

func (basketball *Basketball) startPlay() {
	fmt.Println("Basketball Game Started")
}

func (basketball *Basketball) endPlay() {
	fmt.Println("Basketball Game Finished")
}

type Game struct {
	procedure Procedure
}

// 固定的模板流程
func (game *Game) play() {
	game.procedure.initialize()
	game.procedure.startPlay()
	game.procedure.endPlay()
}

func main() {
	game := &Game{}
	football := &Football{}
	basketball := &Basketball{}

	game.procedure = football
	game.play()

	game.procedure = basketball
	game.play()
}
